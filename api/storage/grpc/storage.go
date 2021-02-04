package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	m "github.com/yuriimakohon/RunecharmsCRUD/api/models"
	"github.com/yuriimakohon/RunecharmsCRUD/pkg/api"
	"google.golang.org/grpc"
	"log"
)

// gRPC implementation of storage.Storage
type Storage struct {
	cli api.CharmCRUDServiceClient
}

func New() *Storage {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &Storage{api.NewCharmCRUDServiceClient(conn)}
}

func (s *Storage) Add(c m.Charm) (m.Charm, error) {
	req := &api.EntityRequest{Entity: CharmModelToGrpc(c)}

	resp, err := s.cli.Add(context.Background(), req)
	if err != nil {
		return m.Charm{}, err
	}

	entity := resp.GetEntities()[0]

	return CharmGrpcToModel(entity), nil
}

func (s *Storage) GetAll() ([]m.Charm, error) {
	resp, err := s.cli.GetAll(context.Background(), &empty.Empty{})
	if err != nil {
		return []m.Charm{}, err
	}

	lenResp, err := s.cli.Len(context.Background(), &empty.Empty{})
	if err != nil {
		return []m.Charm{}, err
	}

	slice := make([]m.Charm, 0, lenResp.Value)
	for _, c := range resp.Entities {
		slice = append(slice, CharmGrpcToModel(c))
	}
	return slice, nil
}

func (s *Storage) Get(id int32) (m.Charm, error) {
	resp, err := s.cli.Get(context.Background(), &api.EntityRequest{Id: id})
	if err != nil {
		return m.Charm{}, err
	}

	return CharmGrpcToModel(resp.Entities[0]), nil
}

func (s *Storage) Delete(id int32) (m.Charm, error) {
	resp, err := s.cli.Delete(context.Background(), &api.EntityRequest{Id: id})
	if err != nil {
		return m.Charm{}, err
	}

	return CharmGrpcToModel(resp.Entities[0]), nil
}

func (s *Storage) Update(id int32, u m.Charm) (m.Charm, error) {
	resp, err := s.cli.Update(context.Background(), &api.EntityRequest{Id: id, Entity: CharmModelToGrpc(u)})
	if err != nil {
		return m.Charm{}, err
	}

	return CharmGrpcToModel(resp.Entities[0]), nil
}

func (s *Storage) Len() (int, error) {
	resp, err := s.cli.Len(context.Background(), &empty.Empty{})
	if err != nil {
		return 0, err
	}
	return int(resp.Value), nil
}
