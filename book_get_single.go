package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// get sinle book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for _, i := range books {
		if i.ID == param["id"] {
			json.NewEncoder(w).Encode(i)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}
