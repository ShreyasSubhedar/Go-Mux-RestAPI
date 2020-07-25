package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// create a book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(len(books) + 1)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}
