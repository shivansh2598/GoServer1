package controllers

import (
	"books-list/model"
	"books-list/repository/book"
	"books-list/utils"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func logFatal(err error){
	if err != nil {
		log.Fatal(err);
	}
}

type Controller struct {}

var books []model.Book
var error2 model.Error

func (c Controller ) GetBooks(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		var book model.Book

		bookRepo := bookRepository.BookRepository{}
		books,err := bookRepo.GetBooks(db, book , books)
		if err != nil {
			error2.Message = "Server Error"
			utils.SendError(w,http.StatusInternalServerError, error2)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w,books);
		return
	}
}


func (c Controller ) GetBook (db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		var book model.Book
		params := mux.Vars(r);
		bookRepo := bookRepository.BookRepository{}
		book, err := bookRepo.GetBook(db,book,books,params)

		if err!=nil {
			error2.Message= "Server Error"
			utils.SendError(w,http.StatusInternalServerError, error2)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w,book)
		return
	}
}

func (c Controller ) AddBook (db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){  // w is the response object and r is the request object
		var book model.Book
		var bookID int
		json.NewDecoder(r.Body).Decode(&book);
		bookRepo := bookRepository.BookRepository{}
		bookID,err := bookRepo.AddBook(db,book,bookID)
		if err != nil {
			error2.Message="Server Error"
			utils.SendError(w, http.StatusInternalServerError, error2)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w,bookID)
		return
	}
}

func (c Controller ) UpdateBook (db *sql.DB ) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		var book model.Book
		json.NewDecoder(r.Body).Decode(&book)
		BookRepo := bookRepository.BookRepository{}
		RowsUpdated,err := BookRepo.UpdateBook(db, book)
		if err !=nil {
			error2.Message="Server Error"
			utils.SendError(w, http.StatusInternalServerError, error2)
		}
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w,RowsUpdated)
		return
	}
}

func (c Controller ) RemoveBook (db *sql.DB ) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		params := mux.Vars(r)
		BookRepo := bookRepository.BookRepository{}
		RowsUpdated,err := BookRepo.RemoveBook(db, params)
		if err != nil {
			error2.Message="Server Error"
			utils.SendError(w,http.StatusInternalServerError,error2)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w,RowsUpdated)
		return
	}
}