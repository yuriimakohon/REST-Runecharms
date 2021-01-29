package main

import (
	"github.com/gorilla/mux"
	"github.com/yuriimakohon/RunecharmsCRUD/api/rest"
	"github.com/yuriimakohon/RunecharmsCRUD/api/storage/grpc"
	"log"
	"net/http"
)

func handleRequest() {
	router := mux.NewRouter().StrictSlash(true)
	s := rest.NewHttpServer(grpc.New())
	// Read
	router.HandleFunc("/charm", s.GetAllCharms).Methods(http.MethodGet)
	router.HandleFunc("/charm/{id}", s.GetCharm).Methods(http.MethodGet)
	// Create
	router.HandleFunc("/charm", s.CreateCharm).Methods(http.MethodPost)
	// Delete
	router.HandleFunc("/charm/{id}", s.DeleteCharm).Methods(http.MethodDelete)
	// Update
	router.HandleFunc("/charm/{id}", s.UpdateCharm).Methods(http.MethodPut)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	handleRequest()
}
