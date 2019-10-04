package controllers

import (
	"books-list/model"
	"books-list/repository/book"
	"books-list/utils"
	"database/sql"
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

func (c Controller ) GetBooks(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		var book model.Book
		var error2 model.Error
		books = []model.Book{}
		bookRepo := bookRepository.BookRepository{}
		books,err := bookRepo.GetBooks(db, book , books)

		if err != nil {
			error2.Message = "Server Error"
			utils.SendError(w,http.StatusInternalServerError, error2)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w,books);

	}
}


//func (c Controller ) GetBook (db *sql.DB) http.HandlerFunc {
//	return func (w http.ResponseWriter, r *http.Request){
//		var book model.Book
//		params := mux.Vars(r);
//		var error2 model.Error
//		bookRepo := bookRepository.BookRepository{}
//		book, err := bookRepo.GetBook(db,book,books,params)
//
//		if err!=nil {
//			error2.Message= "Server Error"
//			utils.SendError(w,http.StatusInternalServerError, error2)
//			return
//		}
//
//		w.Header().Set("Content-Type", "application/json")
//		utils.SendSuccess(w,book)
//	}
//}