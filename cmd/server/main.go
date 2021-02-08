package main

import (
	"github.com/gorilla/mux"
	"github.com/yuriimakohon/RunecharmsCRUD/api/rest"
	"github.com/yuriimakohon/RunecharmsCRUD/api/storage/mongodb"
	"log"
	"net/http"
)

func handleRequest() {
	router := mux.NewRouter().StrictSlash(true)
	s := rest.NewHttpServer(mongodb.New())

	// Read
	router.HandleFunc("/charm", s.GetAllCharms).Methods(http.MethodGet)
	router.HandleFunc("/charm/{id}", s.GetCharm).Methods(http.MethodGet)
	// Create
	router.HandleFunc("/charm", s.CreateCharm).Methods(http.MethodPost)
	// Delete
	router.HandleFunc("/charm/{id}", s.DeleteCharm).Methods(http.MethodDelete)
	// Update
	router.HandleFunc("/charm/{id}", s.UpdateCharm).Methods(http.MethodPut)
	// Len
	router.HandleFunc("/len", s.Len).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	handleRequest()
}
