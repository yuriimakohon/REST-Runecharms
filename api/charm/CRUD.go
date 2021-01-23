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

var charms []Charm

// READ
func GetCharms(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("Endpoint Hit: getCharms")

	json.NewEncoder(w).Encode(charms)
}

func GetCharm(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getCharm")

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println(err)
	} else {
		for _, c := range charms {
			if c.Id == id {
				json.NewEncoder(w).Encode(c)
			}
		}
	}
}

// CREATE
func CreateCharm(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createCharm")

	reqBody, _ := ioutil.ReadAll(r.Body)
	var c Charm
	json.Unmarshal(reqBody, &c)

	charms = append(charms, c)
	json.NewEncoder(w).Encode(c)
}

// DELETE
func DeleteCharm(_ http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteCharm")

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println(err)
	} else {
		for idx, c := range charms {
			if c.Id == id {
				charms = append(charms[:idx], charms[idx+1:]...)
			}
		}
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
		for idx, c := range charms {
			if c.Id == id {
				charms[idx] = u
				json.NewEncoder(w).Encode(charms[idx])
			}
		}
	}
}
