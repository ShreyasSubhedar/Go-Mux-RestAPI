package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book Structs (Model)
type Book struct {
	ID     string  `json:"id"`
	Ispn   string  `json:"ispn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct (Model)
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Mobile    string `json:"mobile"`
}

// Initialize the mock data (book)
var books []Book
var book1 = Book{
	ID:    "1",
	Ispn:  "23423",
	Title: "One night at call center",
	Author: &Author{
		Firstname: "Chetan",
		Lastname:  "Bhagat",
		Mobile:    "9988776655",
	},
}
var book2 = Book{
	ID:    "2",
	Ispn:  "23223",
	Title: "Revolution 2020",
	Author: &Author{
		Firstname: "Chetan",
		Lastname:  "Bhagat",
		Mobile:    "9988776655",
	},
}
var book3 = Book{
	ID:    "3",
	Ispn:  "23223",
	Title: "Five point someone",
	Author: &Author{
		Firstname: "Chetan",
		Lastname:  "Bhagat",
		Mobile:    "9988776655",
	},
}

// get all books
func getBooks(w http.ResponseWriter, r *http.Request) {

}

// get sinle book
func getBook(w http.ResponseWriter, r *http.Request) {

}

// create a book
func createBook(w http.ResponseWriter, r *http.Request) {

}

// update book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

// delete abook
func deleteBook(w http.ResponseWriter, r *http.Request) {

}
func main() {

	fmt.Println("Hello World")
	books = append(books, book1)
	books = append(books, book2)
	books = append(books, book3)

	r := mux.NewRouter()

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}
