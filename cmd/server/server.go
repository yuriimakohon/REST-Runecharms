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
	router.HandleFunc("/charm", charm.GetCharms).Methods("GET")
	router.HandleFunc("/charm/{id}", charm.GetCharm).Methods("GET")
	// Create
	router.HandleFunc("/charm", charm.CreateCharm).Methods("POST")
	// Delete
	router.HandleFunc("/charm/{id}", charm.DeleteCharm).Methods("DELETE")
	// Update
	router.HandleFunc("/charm/{id}", charm.UpdateCharm).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	handleRequest()
}
