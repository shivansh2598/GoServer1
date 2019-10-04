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

func (b BookRepository ) GetBook(db *sql.DB, book model.Book , books []model.Book, params map[string]string ) ( model.Book , error) {

	rows := db.QueryRow("select * from books where id=$1", params["id"])

	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year);

	if err != nil {
		return model.Book{}, err
	}

	return book,nil
}

func (b BookRepository ) AddBook(db *sql.DB, book model.Book, bookID int) ( int,error ) {

	err := db.QueryRow("insert into books (title, author, year) values($1,$2,$3) RETURNING id;",
		book.Title,book.Author, book.Year).Scan(&bookID)

	if err != nil {
		return -1,err;
	}

	return bookID,nil;
}

func (b BookRepository ) UpdateBook(db *sql.DB , book model.Book ) ( int64 , error ){
	result,err := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 RETURNING id", &book.Title, &book.Author,&book.Year,&book.ID);
	if err !=nil {
		return -1, err
	}

	rowsUpdated, err := result.RowsAffected();

	if err != nil {
		return -1,err
	}

	return rowsUpdated,nil
}

func (b BookRepository ) RemoveBook (db *sql.DB , params map[string]string ) (int64 , error){
	result,err:=db.Exec("delete from books where id=$1",params["id"])
	if err != nil {
		return -1,err
	}

	RowsUpdated,err := result.RowsAffected()
	if err != nil {
		return -1, err
	}

	return RowsUpdated,nil
}