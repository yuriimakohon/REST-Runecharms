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

// CREATE
func CreateCharm(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createCharm")

	reqBody, _ := ioutil.ReadAll(r.Body)
	var c Charm
	json.Unmarshal(reqBody, &c)

	if c, err := charms.add(c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		json.NewEncoder(w).Encode(c)
	}
}

// READ
func GetCharms(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getCharms")

	// Filters
	fRune := r.URL.Query().Get("rune")
	fGod := r.URL.Query().Get("god")
	fPower := r.URL.Query().Get("power")

	if charmSlice, err := charms.getAll(); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else if charmSlice != nil {
		filteredSlice := make([]Charm, 0, 1)
		for _, c := range charmSlice {
			if (fRune == "" || c.Rune == fRune) &&
				(fGod == "" || c.God == fGod) &&
				(fPower == "" || fmt.Sprint(c.Power) == fPower) {
				filteredSlice = append(filteredSlice, c)
			}
		}
		json.NewEncoder(w).Encode(filteredSlice)
	}
}

func GetCharm(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getCharm")

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if c, err := charms.get(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(c)
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
	}

	if c, err := charms.update(id, u); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(c)
	}
}

// DELETE
func DeleteCharm(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteCharm")

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, err = charms.delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}
