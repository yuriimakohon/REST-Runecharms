package main

import (
	"github.com/gorilla/mux"
	"github.com/yuriimakohon/RunecharmsCRUD/api/rest"
	"github.com/yuriimakohon/RunecharmsCRUD/api/storage"
	"github.com/yuriimakohon/RunecharmsCRUD/api/storage/elastic"
	"github.com/yuriimakohon/RunecharmsCRUD/api/storage/grpc"
	"github.com/yuriimakohon/RunecharmsCRUD/api/storage/inmem"
	"github.com/yuriimakohon/RunecharmsCRUD/api/storage/mongodb"
	"github.com/yuriimakohon/RunecharmsCRUD/api/storage/postgres"
	"github.com/yuriimakohon/RunecharmsCRUD/api/storage/redis"
	"log"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	var stor storage.Storage
	envStor := os.Getenv("CHARM_STORAGE")

	switch envStor {
	case "redis":
		stor = redis.New()
	case "inmem":
		stor = inmem.New()
	case "postgres":
		stor = postgres.New()
	case "mongo":
		stor = mongodb.New()
	case "grpc":
		stor = grpc.New()
	case "elastic":
		stor = elastic.New()
	default:
		log.Fatal("Invalid 'CHARM_STORAGE' env variable: ", envStor)
		return
	}

	if stor == nil {
		log.Fatal("Storage hasn't created")
		return
	}

	s := rest.NewHttpServer(stor)

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
