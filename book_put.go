package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// update book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for index, item := range books {
		if item.ID == param["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = param["id"]

		}
	}
	json.NewEncoder(w).Encode(books)
}
