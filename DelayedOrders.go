package main

import (
	"fmt"
	//"log"
	//"github.com/gorilla/mux"
	"net/http"
	//"os"
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
)

//отложенный заказ
//суть - заказать файл и ждать пока он выгрузится, проверяем при этом статус
//еще не доделал
type DelayedOrder struct
{
	OrderId string  `db:"order_id"`
	DateStart SimpleDate `db:"date_start"`
	DateEnd SimpleDate  `db:"date_end"`
	Mo string `db:"mo"`
	IsDone null.Bool `db:"is_done"`
	Version string `db:"version"`
	ExportType int `db:"export_type"` //0 через ХП, 1 через код	
}

//сделать заказ на выгрузку 
func createDelayedOrderHandler(w http.ResponseWriter, r *http.Request) {	
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	//get order from request body	
	var err error	
	order := DelayedOrder{}
	porder:= &order
	err = json.NewDecoder(r.Body).Decode(&order)		
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
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
	
	response_order := DelayedOrder{}
	err = tx.Get(&response_order,"INSERT INTO reestr_export.\"order\" (date_start, date_end, mo, is_done) VALUES ( $1, $2, $3, $4) RETURNING order_id", order.DateStart.Time , order.DateEnd.Time, order.Mo, order.IsDone)		
	tx.Commit()

	//get inserted values
	orderid := response_order.OrderId

	err = db.Get(&response_order, "select order_id,date_start,date_end,mo,is_done from reestr_export.\"order\" where order_id = $1;", orderid)

	if err != nil {
        http.Error(w, "Error while getting inserted Order. " + err.Error(), http.StatusInternalServerError)
        return
	}

	//defer GenerateOrderFile(response_order)

 	json.NewEncoder(w).Encode(response_order)	
	return
}

//после создания должна пойти выгрузка
func GenerateDelayedOrderFile (order DelayedOrder) {
	var err error
	_ = err
	fmt.Println("generate")
	//здесь сделать возможным выгрузку по нескольким версиям
	//при выгрузке если происходит ошибка, то отобразить в DelayedOrder
}

//Валидация при разборе json
func (order *DelayedOrder) validate() url.Values {
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