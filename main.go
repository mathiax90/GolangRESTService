package main

import (
	"fmt"
	"log"
	"github.com/gorilla/mux"
	"net/http"
	
	//"strconv"
	//"encoding/json"
	
	//"github.com/fxtlabs/date"
	//"context"	
	"os"
	//pgx "github.com/jackc/pgx/v4"	
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/jackc/pgx/stdlib"
)

func init() {
    // loads values from .env into the system
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
}

var db *sqlx.DB

type CommonError struct {
	Code int
	Text string	
}

func main() {
	//get conf
	DATABASE_URL, exists := os.LookupEnv("DATABASE_URL")
	if !exists {
		fmt.Println("Не найдена строка подключения к БД в файле .env")
	}

	//connect to server (pgx connection no sqlx here)
	// conn, err := pgx.Connect(context.Background(), DATABASE_URL)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	// 	os.Exit(1)
	// }
	// defer conn.Close(context.Background())

	//create db abstraction using pgx driver
	var err error
	db, err = sqlx.Open("pgx", DATABASE_URL )
    if err != nil {
        log.Fatalln(err)
    }	
	defer db.Close()

	XmlCreateTest()
	os.Exit(1)
	
	fmt.Println("service start")
	router := mux.NewRouter()
	router.HandleFunc("/orders/", createOrderHandler).Methods("POST")
	s := http.StripPrefix("/ReestrFileStorage/", http.FileServer(http.Dir("./ReestrFileStorage/")))
	router.PathPrefix("/ReestrFileStorage/").Handler(s)
    log.Fatal(http.ListenAndServe(":8000", router))
}

