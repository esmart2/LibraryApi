package Router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"LibraryApi/Service"
)

func Api(){
	router := mux.NewRouter()

	router.HandleFunc("/addBook", Service.AddBook).Methods("POST")
	router.HandleFunc("/deleteBook/{id}", Service.DeleteBook).Methods("DELETE")
	router.HandleFunc("/getBook/{id}", Service.GetBook).Methods("GET")
	router.HandleFunc("/getAllBooks", Service.GetAllBooks).Methods("GET")
	router.HandleFunc("/editBook/{id}", Service.EditBook).Methods("PUT")

	fmt.Println("Library API is live!!")

	http.ListenAndServe(":8000",router)
}