package Dal

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

type Dal interface {
	AddBook(Book)(string)
	GetBook(string)(Book)
	DeleteBook(string)(string)
	EditBook(string)(string)
	GetAllBooks()(BookList)
}

type Book struct {
	Title string `json: "title"`
	Author string `json: "author"`
	Genre string `json: "genre"`
	Id int `json: "id"`
}

type BookList struct {
	Books []Book `json: "books"`
}

func AddBook(book Book)(string){
	db := dbConnection()

	insert, err := db.Query("INSERT INTO Books VALUES('" + book.Title + "', '" + book.Author + "', '" + book.Genre + "', " + "NULL)")
	
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
	db.Close()

	return "Success"
}

func GetBook(id string)(Book) {
	db := dbConnection()

	result := db.QueryRow("SELECT * FROM Books WHERE id=?;", id)

	var book Book

	err := result.Scan(&book.Author, &book.Title, &book.Genre, &book.Id)
	
	if err != nil {
		panic(err.Error())
	}
	db.Close()

	return book
}

func DeleteBook(id string)(string){
	db := dbConnection()

	db.QueryRow("DELETE FROM Books WHERE id=?;", id)
	db.Close()

	return "Book was succesfully deleted"
}

func EditBook(id string, book Book)(string){
	db := dbConnection()

	str := `update Books set title= '` + book.Title + `', author= '` + book.Author + `', genre= '` + book.Genre + `' where id = ` + id + `;`
	db.Exec(str)
	db.Close()

	return "Successfully updated the book!"
}

func GetAllBooks()(BookList){
	db := dbConnection()

	rows, err := db.Query("SELECT * FROM Books;")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	var book Book
	var list BookList

	for rows.Next() {
		err = rows.Scan(&book.Author, &book.Title, &book.Genre, &book.Id)
		if err != nil {
			panic(err.Error())
		}

		list.Books = append(list.Books, book)
	}
	
	db.Close()

	return list
}

func dbConnection()(*sql.DB){
	// The DB password would never be hard coded here I just did it for convience.
	db, err := sql.Open("mysql", "root:Digicert123@tcp(127.0.0.1:3306)/Library")

	if err != nil {
		panic(err.Error())
	}

	return db
}