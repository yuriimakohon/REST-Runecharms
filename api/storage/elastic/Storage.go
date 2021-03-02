package elastic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	m "github.com/yuriimakohon/RunecharmsCRUD/api/models"
	"github.com/yuriimakohon/RunecharmsCRUD/api/storage"
	"log"
	"reflect"
	"strconv"
)

type Storage struct {
	cli *elastic.Client
}

type lastIdResp struct {
	Value int32 `json:"value"`
}

func New() *Storage {
	var err error
	s := &Storage{}

	s.cli, err = elastic.NewClient()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	s.nextId()
	return s
}

func (s *Storage) nextId() int32 {
	lastId := lastIdResp{Value: 1}
	data, err := json.Marshal(lastId)
	if err != nil {
		return -1
	}

	resp, err := s.cli.Get().Index("seq").Id("lastId").Do(context.Background())
	if err != nil {
		_, _ = s.cli.Index().Index("seq").Id("lastId").BodyJson(string(data)).Do(context.Background())
		return lastId.Value
	}

	if err = json.Unmarshal(resp.Source, &lastId); err != nil {
		return -1
	}
	_, err = s.cli.Update().Index("seq").Id("lastId").Doc(map[string]interface{}{"value": lastId.Value + 1}).Do(context.Background())
	if err != nil {
		return -1
	}

	return lastId.Value
}

func (s *Storage) Add(charm m.Charm) (m.Charm, error) {
	charm.Id = s.nextId()
	data, err := json.Marshal(charm)
	if err != nil {
		return m.Charm{}, err
	}

	resp, err := s.cli.Index().Index("charms").Id(strconv.Itoa(int(charm.Id))).BodyJson(string(data)).Do(context.Background())
	if err != nil {
		return m.Charm{}, err
	}

	intId, err := strconv.Atoi(resp.Id)
	if err != nil {
		return m.Charm{}, err
	}

	charm.Id = int32(intId)
	return charm, nil
}

func (s *Storage) Get(id int32) (m.Charm, error) {
	resp, err := s.cli.Get().Index("charms").Id(strconv.Itoa(int(id))).Do(context.Background())
	if err != nil {
		if elastic.IsNotFound(err) {
			return m.Charm{}, storage.ErrNotFound
		}
		return m.Charm{}, err
	}

	fmt.Print(resp.Found)
	fmt.Print(resp.Error)

	charm := m.Charm{}
	if err = json.Unmarshal(resp.Source, &charm); err != nil {
		return m.Charm{}, err
	}

	return charm, nil
}

func (s *Storage) GetAll() ([]m.Charm, error) {
	capacity, err := s.Len()
	if err != nil || capacity == 0 {
		return []m.Charm{}, err
	}
	charms := make([]m.Charm, 0, capacity)
	charm := m.Charm{}

	searchRes, err := s.cli.Search().Index("charms").Query(elastic.NewMatchAllQuery()).Sort("id", true).Size(capacity).Do(context.Background())
	if err != nil {
		return []m.Charm{}, err
	}

	for _, hit := range searchRes.Each(reflect.TypeOf(charm)) {
		if c, ok := hit.(m.Charm); ok {
			charms = append(charms, c)
		}
	}

	return charms, err
}

func (s *Storage) Delete(id int32) (m.Charm, error) {
	charm, err := s.Get(id)
	if err != nil {
		return m.Charm{}, err
	}

	_, err = s.cli.Delete().Index("charms").Id(strconv.Itoa(int(id))).Do(context.Background())
	if err != nil {
		return m.Charm{}, err
	}

	return charm, nil
}

func (s *Storage) Update(id int32, charm m.Charm) (m.Charm, error) {
	charm.Id = id
	_, err := s.cli.Update().Index("charms").Id(strconv.Itoa(int(id))).Doc(charm).Do(context.Background())
	if elastic.IsNotFound(err) {
		return m.Charm{}, storage.ErrNotFound
	}
	if err != nil {
		return m.Charm{}, err
	}

	return charm, nil
}

func (s *Storage) Len() (int, error) {
	searchRes, err := s.cli.Search().Index("charms").Query(elastic.NewMatchAllQuery()).Pretty(true).Do(context.Background())

	if elastic.IsNotFound(err) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}

	return int(searchRes.Hits.TotalHits.Value), nil
}
