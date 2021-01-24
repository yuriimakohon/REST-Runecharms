package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"runecharm/api/charm"
)

func handleRequest() {
	router := mux.NewRouter().StrictSlash(true)
	// Read
	router.HandleFunc("/charm", charm.GetCharms).Methods(http.MethodGet)
	router.HandleFunc("/charm/{id}", charm.GetCharm).Methods(http.MethodGet)
	// Create
	router.HandleFunc("/charm", charm.CreateCharm).Methods(http.MethodPost)
	// Delete
	router.HandleFunc("/charm/{id}", charm.DeleteCharm).Methods(http.MethodDelete)
	// Update
	router.HandleFunc("/charm/{id}", charm.UpdateCharm).Methods(http.MethodPut)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	handleRequest()
}
