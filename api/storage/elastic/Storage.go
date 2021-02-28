package elastic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	m "github.com/yuriimakohon/RunecharmsCRUD/api/models"
	"log"
)

type Storage struct {
	cli *elastic.Client
}

func New() *Storage {
	var err error
	s := &Storage{}

	s.cli, err = elastic.NewClient()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return s
}

func (s *Storage) Add(charm m.Charm) (m.Charm, error) {
	data, err := json.Marshal(charm)
	if err != nil {
		return m.Charm{}, err
	}

	_, err = s.cli.Index().Index("charms").Id("7").BodyJson(string(data)).Do(context.Background())
	if err != nil {
		return m.Charm{}, err
	}

	return charm, nil
}

func (s *Storage) Get(id int32) (m.Charm, error) {
	resp, err := s.cli.Get().Index("charms").Id("7").Do(context.Background())
	if err != nil {
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
	panic("implement me")
}

func (s *Storage) Delete(id int32) (m.Charm, error) {
	panic("implement me")
}

func (s *Storage) Update(id int32, charm m.Charm) (m.Charm, error) {
	panic("implement me")
}

func (s *Storage) Len() (int, error) {
	panic("implement me")
}
