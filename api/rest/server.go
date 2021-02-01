package rest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	m "github.com/yuriimakohon/RunecharmsCRUD/api/models"
	storage2 "github.com/yuriimakohon/RunecharmsCRUD/api/storage"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Http REST api server that work with internal crud.Storage
type HttpServer struct {
	storage storage2.Storage
}

// Returns new HttpServer with crud.Storage implementation inside
func NewHttpServer(storage storage2.Storage) *HttpServer {
	return &HttpServer{storage: storage}
}

/*
Create new models.Charm in internal crud.Storage

Encode created entity to response
*/
func (s *HttpServer) CreateCharm(w http.ResponseWriter, r *http.Request) {
	// get request body for further unmarshal
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// unmarshal json content from request body to Charm struct
	var c m.Charm
	if err = json.Unmarshal(reqBody, &c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// add created Charm to internal storage
	if c, err = s.storage.Add(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// encode created entity to response
	if err = json.NewEncoder(w).Encode(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/*
Encode all entities from internal crud.Storage to response
*/
func (s *HttpServer) GetAllCharms(w http.ResponseWriter, r *http.Request) {
	// get slice of all Charms from internal storage
	slice, err := s.storage.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// encode slice to response
	if err = json.NewEncoder(w).Encode(slice); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/*
Encode entity with certain id from internal crud.Storage to response
*/
func (s *HttpServer) GetCharm(w http.ResponseWriter, r *http.Request) {
	// get id from request using mux.Vars()
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// get Charm from internal storage by id
	c, err := s.storage.Get(int32(id))
	if err != nil {
		switch err {
		case storage2.ErrNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// encode entity to response
	if err = json.NewEncoder(w).Encode(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/*
Update certain entity with id from internal crud.Storage according to content in json format

Encode updated entity to response
*/
func (s *HttpServer) UpdateCharm(w http.ResponseWriter, r *http.Request) {
	// get request body for further unmarshal
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// unmarshal json content from request body to Charm struct
	var u m.Charm
	if err = json.Unmarshal(reqBody, &u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// get id from request using mux.Vars()
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// update entity in internal storage by id
	c, err := s.storage.Update(int32(id), u)
	if err != nil {
		switch err {
		case storage2.ErrNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// encode updated entity to response
	if err = json.NewEncoder(w).Encode(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/*
Delete certain entity with id from internal crud.Storage
*/
func (s *HttpServer) DeleteCharm(w http.ResponseWriter, r *http.Request) {
	// get id from request using mux.Vars()
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// delete entity from internal storage by id
	if _, err = s.storage.Delete(int32(id)); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

func (s *HttpServer) Len(w http.ResponseWriter, r *http.Request) {
	length, err := s.storage.Len()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err = json.NewEncoder(w).Encode(length); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
