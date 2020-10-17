package Service

import (
	"encoding/json"
	"net/http"
	"LibraryApi/Dal"
	"github.com/gorilla/mux"

)

type Services interface {
	AddBook(http.ResponseWriter, *http.Request) 
	DeleteBook(http.ResponseWriter, *http.Request)
	EditBook(http.ResponseWriter, *http.Request)
	GetBook(http.ResponseWriter, *http.Request)
	GetAllBooks(http.ResponseWriter, *http.Request)
}

type Response struct {
	Message string `json: "message"`
}

func AddBook(w http.ResponseWriter, r *http.Request){
	var newBook Dal.Book
	var resp Response

	err := json.NewDecoder(r.Body).Decode(&newBook)

	if err != nil {
		panic(err.Error())
	}

	if newBook.Title == ""{
		resp.Message = "A book title is required"
	}else if newBook.Author == ""{
		resp.Message = "A author is required"
	}

	resp.Message = Dal.AddBook(newBook)

	json.NewEncoder(w).Encode(resp)

}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	var resp Response
	param := mux.Vars(r)
	id := param["id"]

	resp.Message = Dal.DeleteBook(id)

	json.NewEncoder(w).Encode(resp)
}

func EditBook(w http.ResponseWriter, r *http.Request){
	var book Dal.Book
	var resp Response

	err := json.NewDecoder(r.Body).Decode(&book)
	if err!= nil {
		panic(err.Error())
	}
	param := mux.Vars(r)
	id := param["id"]

	resp.Message = Dal.EditBook(id, book)
	json.NewEncoder(w).Encode(resp)
}

func GetBook(w http.ResponseWriter, r *http.Request){
	var book Dal.Book
	param := mux.Vars(r)
	id := param["id"]

	book = Dal.GetBook(id)

	json.NewEncoder(w).Encode(book)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request){
	var books Dal.BookList

	books = Dal.GetAllBooks()

	json.NewEncoder(w).Encode(books)
}