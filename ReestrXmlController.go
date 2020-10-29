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
	"gopkg.in/guregu/null.v4"
	"net/url"
	"strings"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"bytes"
)

type Order struct
{
	OrderId string `db:"order_id"`
	DateStart SimpleDate `db:"date_start"`
	DateEnd SimpleDate  `db:"date_end"`
	Mo string `db:"mo"`
	IsDone null.Bool `db:"is_done"`
	Version string `db:"version"`
	ExportType int `db:"export_type"` //0 через ХП, 1 через код
	HasError null.Bool `db:"has_error"`
	ErrorText null.String `db:"error_text"`
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
	err = tx.Get(&response_order,"INSERT INTO reestr_export.\"order\" ( date_start, date_end, mo, is_done ) VALUES ( $1, $2, $3, $4) RETURNING order_id", order.DateStart.Time , order.DateEnd.Time, order.Mo, order.IsDone)		

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

func getReestrXml(order_id string) {
	fmt.Println("start getReestrXml")
	var xml_doc string
	var table_fill_ok bool
	var err error

	err = db.Get(&table_fill_ok,"select reestr_export.fn_fill_export_tables($1);",order_id)
	if err != nil {
		fmt.Println("erorr while table fill" + err.Error())
		return
	}

	err = db.Get(&xml_doc,"select reestr_export.sp_get_reestr_xml($1)::text;",order_id)
	if err != nil {
		fmt.Println("erorr while sp_get_reestr_xml" + err.Error())
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

}

// func CreateReestrFile_v3_2_sp(order Order) {
// 	fmt.Println("CreateReestrFile_v3_2_sp");

// 	//start transaction
// 	tx := db.MustBegin()		

// 	var err error
// 	var xml_doc string
// 	err = tx.Get(&xml_doc,"select xml_doc from reestr_export_test('imya faila')")
// 	if err != nil {
// 		fmt.Println("Has error: ", err);
// 		tx.Rollback()
// 	}
// 	//write to file here

// 	//xml ok, update 
// 	_, err = db.NamedExec(`update "order" set is_done = false where order_id = :order_id`, 
//         map[string]interface{}{
//             "order_id": order.OrderId,           
// 	})
// 	if err != nil {
// 		fmt.Println("Has error on update row: ", err);
// 		tx.Rollback()
// 	}

// 	tx.Commit()	
// 	return
// }

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