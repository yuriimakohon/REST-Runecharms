package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/yuriimakohon/RunecharmsCRUD/api/storage"
	"github.com/yuriimakohon/RunecharmsCRUD/api/storage/grpc"
	"github.com/yuriimakohon/RunecharmsCRUD/pkg/api"
)

type charmCRUDService struct {
	storage storage.Storage
}

func New(storage storage.Storage) *charmCRUDService {
	return &charmCRUDService{storage: storage}
}

func (s charmCRUDService) Add(ctx context.Context, request *api.EntityRequest) (*api.EntityResponse, error) {
	mCharm := grpc.CharmGrpcToModel(request.Entity)

	resp := &api.EntityResponse{
		Entities: make([]*api.Charm, 0, 1),
	}

	mCharm, err := s.storage.Add(mCharm)
	if err != nil {
		return resp, err
	}

	apiCharm := grpc.CharmModelToGrpc(mCharm)

	resp.Entities = append(resp.Entities, apiCharm)
	return resp, nil
}

func (s charmCRUDService) GetAll(ctx context.Context, empty *empty.Empty) (*api.EntityResponse, error) {
	length, err := s.storage.Len()
	if err != nil {
		return nil, err
	}

	resp := &api.EntityResponse{
		Entities: make([]*api.Charm, 0, length),
	}

	charms, err := s.storage.GetAll()
	if err != nil {
		return nil, err
	}

	for _, c := range charms {
		resp.Entities = append(resp.Entities, grpc.CharmModelToGrpc(c))
	}
	return resp, nil
}

func (s charmCRUDService) Get(ctx context.Context, request *api.EntityRequest) (*api.EntityResponse, error) {
	mCharm, err := s.storage.Get(request.Id)
	if err != nil {
		return nil, err
	}

	apiCharm := grpc.CharmModelToGrpc(mCharm)

	resp := &api.EntityResponse{Entities: append(make([]*api.Charm, 0, 1), apiCharm)}
	return resp, nil
}

func (s charmCRUDService) Delete(ctx context.Context, request *api.EntityRequest) (*api.EntityResponse, error) {
	mCharm, err := s.storage.Delete(request.Id)
	if err != nil {
		return nil, err
	}

	apiCharm := grpc.CharmModelToGrpc(mCharm)

	resp := &api.EntityResponse{Entities: append(make([]*api.Charm, 0, 1), apiCharm)}
	return resp, nil
}

func (s charmCRUDService) Update(ctx context.Context, request *api.EntityRequest) (*api.EntityResponse, error) {
	mCharm, err := s.storage.Update(request.Id, grpc.CharmGrpcToModel(request.Entity))
	if err != nil {
		return nil, err
	}

	apiCharm := grpc.CharmModelToGrpc(mCharm)

	resp := &api.EntityResponse{Entities: append(make([]*api.Charm, 0, 1), apiCharm)}
	return resp, nil
}

func (s charmCRUDService) Len(ctx context.Context, empty *empty.Empty) (*api.LenResponse, error) {
	length, err := s.storage.Len()
	if err != nil {
		return nil, err
	}

	return &api.LenResponse{Value: int32(length)}, nil
}
