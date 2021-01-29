package storage

import (
	"errors"
	m "github.com/yuriimakohon/RunecharmsCRUD/api/models"
)

var (
	ErrNotFound = errors.New("charm not found")
	_           = ErrNotFound
)

type Storage interface {
	Add(charm m.Charm) (m.Charm, error)
	Get(id int32) (m.Charm, error)
	GetAll() ([]m.Charm, error)
	Delete(id int32) (m.Charm, error)
	Update(id int32, charm m.Charm) (m.Charm, error)
}
