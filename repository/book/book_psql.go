package bookRepository

import (
	"books-list/model"
	"database/sql"
)

type BookRepository struct{}

//psql query result

func (b BookRepository) GetBooks(db *sql.DB, book model.Book , books []model.Book) ([]model.Book, error) {
	rows,err := db.Query("select * from books")

	if err != nil {
		return []model.Book{}, err
	}

	defer rows.Close();

	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		books = append(books, book)
	}

	if err != nil {
		return []model.Book{},err
	}

	return books, nil
}

//func (b BookRepository ) GetBook(db *sql.DB, book model.Book , books []model.Book, params map[string]string ) ( model.Book , error) {
//
//	rows := db.QueryRow("select * from books where id=$1", params["id"])
//
//	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year);
//
//	if err != nil {
//		return model.Book{}, err
//	}
//
//	return book,nil
//}
