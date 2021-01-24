package charm

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var charms Storage

func init() {
	charms = &IMCharms{
		storage: make([]Charm, 0, 1),
		lastId:  0,
	}
}

// READ
func GetCharms(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getCharms")

	charmSlice := charms.getAll()
	if charmSlice != nil {
		json.NewEncoder(w).Encode(charmSlice)
	}
}

func GetCharm(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getCharm")

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println(err)
	} else if c := charms.get(id); c != nil {
		json.NewEncoder(w).Encode(c)
	}
}

// CREATE
func CreateCharm(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createCharm")

	reqBody, _ := ioutil.ReadAll(r.Body)
	var c Charm
	json.Unmarshal(reqBody, &c)

	json.NewEncoder(w).Encode(charms.add(c))
}

// DELETE
func DeleteCharm(_ http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteCharm")

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println(err)
	} else {
		charms.delete(id)
	}
}

// UPDATE
func UpdateCharm(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: updateCharm")

	reqBody, _ := ioutil.ReadAll(r.Body)
	var u Charm
	json.Unmarshal(reqBody, &u)

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println(err)
	} else {
		json.NewEncoder(w).Encode(charms.update(id, u))
	}
}
