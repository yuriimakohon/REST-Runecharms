package mongodb

import (
	"context"
	m "github.com/yuriimakohon/RunecharmsCRUD/api/models"
	"github.com/yuriimakohon/RunecharmsCRUD/api/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Storage struct {
	client     *mongo.Client
	coll       *mongo.Collection
	ctx        context.Context
	lastIdColl *mongo.Collection
}

func New() *Storage {
	st := &Storage{ctx: context.TODO()}
	var err error

	st.client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
		return nil
	}

	// Create connect
	err = st.client.Connect(st.ctx)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	// Check the connection
	err = st.client.Ping(st.ctx, nil)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	// Check is lastId already exist
	if err = st.createLastId(); err != nil {
		log.Fatal(err)
		return nil
	}

	st.coll = st.client.Database("test").Collection("charms")
	return st
}

func (s *Storage) createLastId() error {
	s.lastIdColl = s.client.Database("test").Collection("seq")

	opRes, err := s.lastIdColl.CountDocuments(s.ctx, bson.D{})
	if err != nil || opRes == 0 {
		_, err = s.lastIdColl.InsertOne(s.ctx, bson.D{
			{"type", "charm_id"},
			{"lastId", 0},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Storage) nextLastId() int32 {
	filter := bson.D{{"type", "charm_id"}}
	update := bson.D{
		{"$inc", bson.D{
			{"lastId", 1},
		}},
	}
	res := bson.D{}

	opRes := s.lastIdColl.FindOne(s.ctx, filter)

	if err := opRes.Decode(&res); err != nil {
		log.Fatal(err)
		return -1
	}

	if _, err := s.lastIdColl.UpdateOne(s.ctx, filter, update); err != nil {
		log.Fatal(err)
		return -1
	}

	return res.Map()["lastId"].(int32)
}

func (s *Storage) Add(charm m.Charm) (m.Charm, error) {
	charm.Id = s.nextLastId()
	_, err := s.coll.InsertOne(s.ctx, charm)
	if err != nil {
		return m.Charm{}, err
	}

	return charm, nil
}

func (s *Storage) Get(id int32) (m.Charm, error) {
	filter := bson.D{{"id", id}}
	charm := m.Charm{}

	opRes := s.coll.FindOne(s.ctx, filter)

	if err := opRes.Decode(&charm); err != nil {
		return m.Charm{}, storage.ErrNotFound
	}
	return charm, nil
}

func (s *Storage) GetAll() ([]m.Charm, error) {
	filter := bson.M{}

	l, err := s.Len()
	if err != nil {
		return []m.Charm{}, nil
	}

	charms := make([]m.Charm, 0, l)

	cur, err := s.coll.Find(s.ctx, filter)
	if err != nil {
		return []m.Charm{}, err
	}

	for cur.Next(s.ctx) {
		var el m.Charm
		err = cur.Decode(&el)
		if err != nil {
			log.Fatal(err)
		}

		charms = append(charms, el)
	}
	return charms, nil
}

func (s *Storage) Delete(id int32) (m.Charm, error) {
	filter := bson.D{{"id", id}}

	charm, err := s.Get(id)
	if err != nil {
		return m.Charm{}, err
	}

	_, err = s.coll.DeleteOne(s.ctx, filter)
	if err != nil {
		return m.Charm{}, err
	}

	return charm, nil
}

func (s *Storage) Update(id int32, charm m.Charm) (m.Charm, error) {
	filter := bson.D{{"id", id}}
	update := bson.D{
		{"$set", bson.D{
			{"rune", charm.Rune},
			{"god", charm.God},
			{"power", charm.Power},
		}},
	}

	_, err := s.coll.UpdateOne(s.ctx, filter, update)
	if err != nil {
		return m.Charm{}, err
	}

	charm.Id = id
	return charm, nil
}

func (s *Storage) Len() (int, error) {
	opRes, err := s.coll.CountDocuments(s.ctx, bson.D{})
	if err != nil {
		return 0, err
	}

	return int(opRes), err
}
