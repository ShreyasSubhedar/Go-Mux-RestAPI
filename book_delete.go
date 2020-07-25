package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// delete abook
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for index, item := range books {
		if item.ID == param["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}
