package main

import (
	"books-list/controllers"
	"books-list/driver"
	"books-list/model"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
)

//type Book struct {
//	ID int `json:id`
//	Title string `json:title`
//	Author string `json:author`
//	Year string `json:year`
//}

var books []model.Book
var db *sql.DB

func init(){
	gotenv.Load();
}

func logFatal(err error){
	if err != nil {
		log.Fatal(err);
	}
}

func main(){

	db = driver.ConnectDB();
 	controller := controllers.Controller{}

	router := mux.NewRouter();

	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.RemoveBook(db)).Methods("DELETE")

	fmt.Println("Server is runnig at port 8000");
	log.Fatal(http.ListenAndServe(":8000", router))
}


