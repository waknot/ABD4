package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", NewReservation).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func NewReservation(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("reservation ok")
}
