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
	"github.com/rs/cors"
	"strings"
	"io"
	"bytes"
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

	//XmlCreateTest()
	//os.Exit(1)
	
	fmt.Println("service start")
	router := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},                            // All origins
		AllowedMethods: []string{"POST", "GET", "PUT", "DELETE"}, // Allowing only get, just an example
	})

	router.HandleFunc("/orders/", createOrderHandler).Methods("POST")
	router.HandleFunc("/orders/", getOrderHandler).Methods("GET")
	router.HandleFunc("/orders/idpacs", getIdPacsHandler).Methods("GET")
	router.HandleFunc("/files/upload", UploadFile).Methods("POST")
	s := http.StripPrefix("/ReestrFileStorage/", http.FileServer(http.Dir("./ReestrFileStorage/")))
	router.PathPrefix("/ReestrFileStorage/").Handler(s)
    log.Fatal(http.ListenAndServe(":8000", c.Handler(router)))
}

func UploadFile(w http.ResponseWriter, r *http.Request){
	r.ParseMultipartForm(32 << 20) // limit your max input length!
	var buf bytes.Buffer	
    // in your case file would be fileupload
    file, header, err := r.FormFile("file1")
    if err != nil {
        panic(err)
    }
    defer file.Close()
    name := strings.Split(header.Filename, ".")
    fmt.Printf("File name %s\n", name[0])
    // Copy the file data to my buffer
    io.Copy(&buf, file)
    // do something with the contents...
    // I normally have a struct defined and unmarshal into a struct, but this will
    // work as an example
    contents := buf.String()
    fmt.Println(contents)
    // I reset the buffer in case I want to use it again
    // reduces memory allocations in more intense projects
    buf.Reset()
    // do something else
    // etc write header
    return
}
