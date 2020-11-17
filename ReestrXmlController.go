package main

import (
	"fmt"
	//"log"
	//"github.com/gorilla/mux"
	"net/http"
	"os"
	//"strconv"
	"encoding/json"
	//"math/rand"
	//"github.com/fxtlabs/date"
	//"github.com/jackc/pgx"
	//"github.com/jmoiron/sqlx"
	//"database/sql"
	"time"
	//"gopkg.in/guregu/null.v4"
	"net/url"
	"strings"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"bytes"
)


type ID_PAC struct
{
	IdPac string `db:"ID_PAC"`
}

type Order struct
{
	OrderId string `db:"order_id"`
	DateStart SimpleDate `db:"date_start"`
	DateEnd SimpleDate  `db:"date_end"`
	Mo string `db:"mo"`
	State int `db:"state"`
	// Version string `db:"version"`
	// ExportType int `db:"export_type"` //0 через ХП, 1 через код
	// HasError null.Bool `db:"has_error"`
	// ErrorText null.String `db:"error_text"`
}


func createOrderHandler(w http.ResponseWriter, r *http.Request) {	
	fmt.Println("crate order and file")
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	//get order from request body	
	var err error	
	order := Order{}
	porder:= &order
	err = json.NewDecoder(r.Body).Decode(&order)		
    if err != nil {
        http.Error(w, "Json obj decode err.\n"+ err.Error(), http.StatusBadRequest)
        return
	}

	//validate
	if validErrs := porder.validate(); len(validErrs) > 0 {
		err := map[string]interface{}{"validationError": validErrs}		
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	
	//insert into db
	tx := db.MustBegin()		
	
	response_order := Order{}
	err = tx.Get(&response_order,"INSERT INTO reestr_export.\"order\" ( date_start, date_end, mo, state ) VALUES ( $1, $2, $3, $4) RETURNING order_id", order.DateStart.Time , order.DateEnd.Time, order.Mo, 0)		

	if err != nil {
		tx.Rollback()
		http.Error(w, "Table Order Insert Error.\n"+ err.Error(), http.StatusInternalServerError)	
		return
	}	

	tx.Commit()
	defer getReestrXml(response_order.OrderId)
	//fmt.Println(response_order)

	 if err = json.NewEncoder(w).Encode(response_order.OrderId); err != nil {
		 fmt.Println(err)
		 http.Error(w, "Error on struct encode to json.\n"+ err.Error(), http.StatusInternalServerError)	
	}
}

func getOrderHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var err error
	fmt.Println("get order status")
	var order_id string
	if len(r.URL.RawQuery) > 0 {
		order_id = r.URL.Query().Get("order_id")
		//fmt.Printf(str1)
		if order_id == "" {
			http.Error(w, "order_id is not defined\n", http.StatusBadRequest)
			return
		}
	}

	fmt.Println("order_id: " + order_id)
	order:= Order{}
	err = db.Get(&order,`select order_id, date_start, date_end, mo, state from reestr_export."order" where order_id = $1`, order_id)
	
	fmt.Println(order)

	if err != nil {		
		http.Error(w, "Error while get order info ("+order_id+").\n"+ err.Error(), http.StatusInternalServerError)	
		return
	}	
	
	if err = json.NewEncoder(w).Encode(order); err != nil {
		http.Error(w, "Error while Unmarshal order\n"+ err.Error(), http.StatusInternalServerError)	
	}
	
}

func getReestrXml(order_id string) {
	fmt.Println("start getReestrXml")
	var xml_doc string
	var table_fill_ok bool
	var err error

	err = db.Get(&table_fill_ok,"select reestr_export.fn_fill_export_tables($1);",order_id)
	if err != nil {
		fmt.Println("erorr while table fill" + err.Error())
		_, err = db.Exec(`update reestr_export."order" set state = 2 where order_id = $1`,order_id)
		return
	}

	err = db.Get(&xml_doc,"select reestr_export.sp_get_reestr_xml($1)::text;",order_id)
	if err != nil {
		fmt.Println("erorr while sp_get_reestr_xml" + err.Error())
		_, err = db.Exec(`update reestr_export."order" set state = 2 where order_id = $1`,order_id)
		return
	}

	xml_doc = `<?xml version="1.0" encoding="windows-1251"?>` + xml_doc
	//Запись в файл
	file, err := os.Create("ReestrFileStorage/" + order_id +".xml")
     
    if err != nil{
		fmt.Println("erorr while file create")
		return	
	}

	defer file.Close()

	var win1251Bytes bytes.Buffer
	win1251Transform := transform.NewWriter(&win1251Bytes, charmap.Windows1251.NewEncoder())
	win1251Transform.Write([]byte(xml_doc))
	win1251Transform.Close()
	file.WriteString(win1251Bytes.String())

	_, err = db.Exec(`update reestr_export."order" set state = 1 where order_id = $1`,order_id)
}

func (order *Order) validate() url.Values {
	errs := url.Values{}

	// check if the title empty
	if strings.Trim(order.Mo," ") == "" {
		errs.Add("Mo", "Mo is required")
	}

	// check the title field is between 3 to 120 chars
	if !order.DateStart.Time.Before(order.DateEnd.Time){
		errs.Add("DateStart", "Is greater than DateEnd")
	}

	minDate, _ := time.Parse("2006-01-02", "2010-01-01")
	if order.DateStart.Time.Before(minDate){
		errs.Add("DateStart", "Is less than 01.01.2010")
	}

	if order.DateEnd.Time.Before(minDate){
		errs.Add("DateEnd", "Is less than 01.01.2010")
	}
	return errs
}

func getIdPacsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var err error
	fmt.Println("getIdPacs")
	var order_id string
	if len(r.URL.RawQuery) > 0 {
		order_id = r.URL.Query().Get("order_id")
		//fmt.Printf(str1)
		if order_id == "" {
			http.Error(w, "order_id is not defined\n", http.StatusBadRequest)
			return
		}
	}

	fmt.Println("order_id: " + order_id)
	var ID_PACs []ID_PAC
	
	err = db.Select(&ID_PACs,"select * from reestr_export.sp_get_id_pacs($1)", order_id)		

	if err != nil {		
		http.Error(w, "Error while  reestr_export.sp_get_id_pacs("+order_id+").\n"+ err.Error(), http.StatusInternalServerError)	
		return
	}	
	
	if err = json.NewEncoder(w).Encode(ID_PACs); err != nil {
		http.Error(w, "Error while Unmarshal ID_PACs\n"+ err.Error(), http.StatusInternalServerError)	
	}
	
}
	