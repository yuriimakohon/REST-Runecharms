package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	m "github.com/yuriimakohon/RunecharmsCRUD/api/models"
	"github.com/yuriimakohon/RunecharmsCRUD/api/storage"
	"log"
	"strconv"
)

type Storage struct {
	ctx context.Context
	rdb *redis.Client
}

func New() *Storage {
	st := &Storage{
		context.Background(),
		redis.NewClient(&redis.Options{Addr: "localhost:6379"}),
	}

	isExists, err := st.rdb.Exists(st.ctx, "charm.lastId").Result()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	if isExists == 0 {
		st.rdb.Set(st.ctx, "charm.lastId", 0, 0)
	}
	return st
}

func (s *Storage) nextId() int32 {
	id, _ := s.rdb.Incr(s.ctx, "charm.lastId").Result()
	return int32(id)
}

func (s *Storage) Add(charm m.Charm) (m.Charm, error) {
	id := s.nextId()
	s.rdb.Set(s.ctx, sRune(id), charm.Rune, 0)
	s.rdb.Set(s.ctx, sGod(id), charm.God, 0)
	s.rdb.Set(s.ctx, sPower(id), charm.Power, 0)
	s.rdb.SAdd(s.ctx, "idSet", id)

	charm.Id = id
	return charm, nil
}

func (s *Storage) Get(id int32) (m.Charm, error) {
	charm := m.Charm{Id: id}
	var err error

	if charm.Rune, err = s.rdb.Get(s.ctx, sRune(id)).Result(); err != nil {
		return m.Charm{}, storage.ErrNotFound
	}

	if charm.God, err = s.rdb.Get(s.ctx, sGod(id)).Result(); err != nil {
		return m.Charm{}, storage.ErrNotFound
	}

	strPower, err := s.rdb.Get(s.ctx, sPower(id)).Result()
	if err != nil {
		return m.Charm{}, storage.ErrNotFound
	}
	intPower, err := strconv.Atoi(strPower)
	if err != nil {
		return m.Charm{}, err
	}
	charm.Power = int32(intPower)

	return charm, nil
}

func (s *Storage) GetAll() ([]m.Charm, error) {
	charm := m.Charm{}
	charms := make([]m.Charm, 0, 1)

	idSet, err := s.rdb.SMembers(s.ctx, "idSet").Result()
	if err != nil {
		return []m.Charm{}, err
	}

	for _, strId := range idSet {
		intId, err := strconv.Atoi(strId)
		if err != nil {
			return []m.Charm{}, err
		}

		charm, err = s.Get(int32(intId))
		if err != nil {
			return []m.Charm{}, err
		}

		charms = append(charms, charm)
	}

	return charms, nil
}

func (s *Storage) Delete(id int32) (m.Charm, error) {
	charm, err := s.Get(id)
	if err != nil {
		return m.Charm{}, storage.ErrNotFound
	}

	s.rdb.Del(s.ctx, sRune(id))
	s.rdb.Del(s.ctx, sGod(id))
	s.rdb.Del(s.ctx, sPower(id))
	s.rdb.SRem(s.ctx, "idSet", id)

	return charm, nil
}

func (s *Storage) Update(id int32, charm m.Charm) (m.Charm, error) {
	_, err := s.Get(id)
	if err != nil {
		return m.Charm{}, storage.ErrNotFound
	}

	s.rdb.Set(s.ctx, sRune(id), charm.Rune, 0)
	s.rdb.Set(s.ctx, sGod(id), charm.God, 0)
	s.rdb.Set(s.ctx, sPower(id), charm.Power, 0)

	charm.Id = id
	return charm, nil
}

func (s *Storage) Len() (int, error) {
	length, err := s.rdb.SCard(s.ctx, "idSet").Result()
	if err != nil {
		return 0, err
	}

	return int(length), nil
}
