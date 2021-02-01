package main

import (
	mygrpc "github.com/yuriimakohon/RunecharmsCRUD/api/grpc"
	"github.com/yuriimakohon/RunecharmsCRUD/api/storage/inmem"
	"github.com/yuriimakohon/RunecharmsCRUD/pkg/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv := mygrpc.New(inmem.New())
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
