package main

import (
	"books-list/controllers"
	"books-list/driver"
	"books-list/model"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
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

	//cors enabled code snippet
	headers :=  handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string {"*"})


	//router method
	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.RemoveBook(db)).Methods("DELETE")
	fmt.Println("Server is running at port 8000");
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(headers,methods,origins)(router))) //server listens at port 8000

}


