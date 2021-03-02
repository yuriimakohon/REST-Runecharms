package main

import (
	mygrpc "github.com/yuriimakohon/RunecharmsCRUD/api/grpc"
	"github.com/yuriimakohon/RunecharmsCRUD/api/storage"
	"github.com/yuriimakohon/RunecharmsCRUD/api/storage/elastic"
	"github.com/yuriimakohon/RunecharmsCRUD/api/storage/inmem"
	"github.com/yuriimakohon/RunecharmsCRUD/api/storage/mongodb"
	"github.com/yuriimakohon/RunecharmsCRUD/api/storage/postgres"
	"github.com/yuriimakohon/RunecharmsCRUD/api/storage/redis"
	"github.com/yuriimakohon/RunecharmsCRUD/pkg/api"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	s := grpc.NewServer()
	var stor storage.Storage
	envStor := os.Getenv("CHARM_GRPC_STORAGE")

	switch envStor {
	case "redis":
		stor = redis.New()
	case "inmem":
		stor = inmem.New()
	case "postgres":
		stor = postgres.New()
	case "mongo":
		stor = mongodb.New()
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

	srv := mygrpc.New(stor)
	api.RegisterCharmCRUDServiceServer(s, srv)

	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
		return
	}

	if err = s.Serve(l); err != nil {
		log.Fatal(err)
		return
	}
}
