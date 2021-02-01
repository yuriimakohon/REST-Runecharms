package inmem

import (
	m "github.com/yuriimakohon/RunecharmsCRUD/api/models"
	"github.com/yuriimakohon/RunecharmsCRUD/api/storage"
)

// In-memory implementation of storage.Storage
type Storage struct {
	Charms []m.Charm
	LastId int32
}

func New() *Storage {
	return &Storage{make([]m.Charm, 0, 10), 0}
}

func (s *Storage) Add(c m.Charm) (m.Charm, error) {
	c.Id = s.LastId
	s.Charms = append(s.Charms, c)
	s.LastId++
	return c, nil
}

func (s *Storage) GetAll() ([]m.Charm, error) {
	cp := make([]m.Charm, len(s.Charms))
	copy(cp, s.Charms)
	return cp, nil
}

func (s *Storage) Get(id int32) (m.Charm, error) {
	for _, c := range s.Charms {
		if c.Id == id {
			return c, nil
		}
	}
	return m.Charm{}, storage.ErrNotFound
}

func (s *Storage) Delete(id int32) (m.Charm, error) {
	for idx, c := range s.Charms {
		if c.Id == id {
			s.Charms = append(s.Charms[:idx], s.Charms[idx+1:]...)
			return c, nil
		}
	}
	return m.Charm{}, storage.ErrNotFound
}

func (s *Storage) Update(id int32, u m.Charm) (m.Charm, error) {
	for idx, c := range s.Charms {
		if c.Id == id {
			u.Id = id
			s.Charms[idx] = u
			return s.Charms[idx], nil
		}
	}
	return m.Charm{}, storage.ErrNotFound
}

func (s *Storage) Len() (int, error) {
	return len(s.Charms), nil
}
