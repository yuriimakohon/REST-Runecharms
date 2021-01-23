package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"runecharm/api/charm"
)

func handleRequest() {
	gRouter := mux.NewRouter().StrictSlash(true)
	// Read
	gRouter.HandleFunc("/charm", charm.GetCharms).Methods("GET")
	gRouter.HandleFunc("/charm/{id}", charm.GetCharm).Methods("GET")
	// Create
	gRouter.HandleFunc("/charm", charm.CreateCharm).Methods("POST")
	// Delete
	gRouter.HandleFunc("/charm/{id}", charm.DeleteCharm).Methods("DELETE")
	// Update
	gRouter.HandleFunc("/charm/{id}", charm.UpdateCharm).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", gRouter))
}

func main() {
	handleRequest()
}
