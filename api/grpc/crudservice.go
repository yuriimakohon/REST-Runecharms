package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	m "github.com/yuriimakohon/RunecharmsCRUD/api/models"
	"github.com/yuriimakohon/RunecharmsCRUD/api/storage"
	"github.com/yuriimakohon/RunecharmsCRUD/pkg/api"
)

type charmCRUDService struct {
	storage storage.Storage
}

func New(storage storage.Storage) *charmCRUDService {
	return &charmCRUDService{storage: storage}
}

func (c charmCRUDService) Add(ctx context.Context, request *api.EntityRequest) (*api.EntityResponse, error) {
	mCharm := m.Charm{
		Rune:  request.Entity.Rune,
		God:   request.Entity.God,
		Power: request.Entity.Power,
	}

	resp := &api.EntityResponse{
		Entities: make([]*api.Charm, 0, 1),
	}

	mCharm, err := c.storage.Add(mCharm)
	if err != nil {
		return resp, err
	}

	apiCharm := &api.Charm{
		Id:    mCharm.Id,
		Rune:  mCharm.Rune,
		God:   mCharm.God,
		Power: mCharm.Power,
	}

	resp.Entities = append(resp.Entities, apiCharm)
	return resp, nil
}

func (c charmCRUDService) GetAll(ctx context.Context, empty *empty.Empty) (*api.EntityResponse, error) {
	panic("implement me")
}

func (c charmCRUDService) Get(ctx context.Context, request *api.EntityRequest) (*api.EntityResponse, error) {
	panic("implement me")
}

func (c charmCRUDService) Delete(ctx context.Context, request *api.EntityRequest) (*api.EntityResponse, error) {
	panic("implement me")
}

func (c charmCRUDService) Update(ctx context.Context, request *api.EntityRequest) (*api.EntityResponse, error) {
	panic("implement me")
}
