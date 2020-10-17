package Router

import (
	"bytes"
	"testing"
	"net/http/httptest"
	"net/http"
	"LibraryApi/Service"
	// "github.com/gorilla/mux"
)

func TestAddBook(t *testing.T){
	var body = []byte(`{"author": "CS Lewis", "title": "Narnia", "genre": "fiction"}`)
	
	req, err := http.NewRequest("POST", "/addBook", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Service.AddBook)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != 200 {
		t.Errorf("Test failed to add book to the library")
	}
}

func TestGetAllBooks(t *testing.T){
	req, err := http.NewRequest("GET", "/getAllBooks", nil)
	
	if err != nil {
		t.Errorf("Trouble creating request for  GetAllBooks")
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Service.GetAllBooks)
	handler.ServeHTTP(rr,req)

	if rr.Code != 200 {
		t.Errorf("Expected response code 200")
	}
}

// func TestGetBook(t *testing.T){
// 	req, err := http.NewRequest("GET", "/getBook/1", nil)

// 	if err != nil {
// 		t.Errorf("Failed creating http request")
// 	}

// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(Service.GetBook)
// 	handler.ServeHTTP(rr,req)

// 	if rr.Code != 200 {
// 		t.Errorf("Expected response code 200")
// 	}
// }

func TestEditBook(t *testing.T){
	var body = []byte(`{"author": "C.S. Lewis", "title": "The Chronicles Narnia", "genre": "fiction"}`)
	
	req, err := http.NewRequest("PUT", "/editBook/1", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Service.AddBook)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != 200 {
		t.Errorf("Test failed to add book to the library")
	}
}

func TestDeleteBook(t *testing.T){	
	req, err := http.NewRequest("DELETE", "/deleteBook/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Service.DeleteBook)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != 200 {
		t.Errorf("Test failed to add book to the library")
	}
}