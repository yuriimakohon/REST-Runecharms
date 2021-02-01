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

func (s charmCRUDService) Add(ctx context.Context, request *api.EntityRequest) (*api.EntityResponse, error) {
	mCharm := m.Charm{
		Rune:  request.Entity.Rune,
		God:   request.Entity.God,
		Power: request.Entity.Power,
	}

	resp := &api.EntityResponse{
		Entities: make([]*api.Charm, 0, 1),
	}

	mCharm, err := s.storage.Add(mCharm)
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
		resp.Entities = append(
			resp.Entities,
			&api.Charm{
				Id:    c.Id,
				Rune:  c.Rune,
				God:   c.God,
				Power: c.Power,
			},
		)
	}
	return resp, nil
}

func (s charmCRUDService) Get(ctx context.Context, request *api.EntityRequest) (*api.EntityResponse, error) {
	panic("implement me")
}

func (s charmCRUDService) Delete(ctx context.Context, request *api.EntityRequest) (*api.EntityResponse, error) {
	panic("implement me")
}

func (s charmCRUDService) Update(ctx context.Context, request *api.EntityRequest) (*api.EntityResponse, error) {
	panic("implement me")
}

func (s charmCRUDService) Len(ctx context.Context, empty *empty.Empty) (*api.LenResponse, error) {
	length, err := s.storage.Len()
	if err != nil {
		return nil, err
	}

	return &api.LenResponse{Value: int32(length)}, nil
}
