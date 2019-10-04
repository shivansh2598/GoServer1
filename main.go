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

//func getBooks(w http.ResponseWriter, r *http.Request){
//	var book model.Book
//	books = []model.Book{}
//
//	rows,err := db.Query("select * from books")
//	logFatal(err)
//
//	defer rows.Close();
//
//	for rows.Next() {
//		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year);
//		logFatal(err);
//
//		books = append(books, book);
//	}
//
//	json.NewEncoder(w).Encode(books);
//}

//func getBook(w http.ResponseWriter, r *http.Request){
//	var book model.Book
//	params := mux.Vars(r);
//
//	rows := db.QueryRow("select * from books where id=$1", params["id"])
//
//	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year);
//	logFatal(err);
//
//	json.NewEncoder(w).Encode(book)
//
//}

//func addBook(w http.ResponseWriter, r *http.Request){  // w is the response object and r is the request object
//	var book model.Book
//	var bookID int
//
//	json.NewDecoder(r.Body).Decode(&book);
//
//	err := db.QueryRow("insert into books (title, author, year) values($1,$2,$3) RETURNING id;",
//		book.Title,book.Author, book.Year).Scan(&bookID)
//
//	logFatal(err);
//
//	json.NewEncoder(w).Encode(bookID)
//
//}

//func updateBook(w http.ResponseWriter, r *http.Request){
//	var book model.Book
//	json.NewDecoder(r.Body).Decode(&book)
//
//	result,err := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 RETURNING id", &book.Title, &book.Author,&book.Year,&book.ID);
//	logFatal(err)
//
//	rowsUpdated, err := result.RowsAffected();
//	logFatal(err);
//
//	json.NewEncoder(w).Encode(rowsUpdated);
//}

//func removeBook(w http.ResponseWriter, r *http.Request){
//	params := mux.Vars(r);
//
//	result, err := db.Exec("delete from books where id = $1", params["id"])
//	logFatal(err)
//
//	rowsUpdated,err := result.RowsAffected();
//	logFatal(err);
//
//	json.NewEncoder(w).Encode(rowsUpdated);
//}

