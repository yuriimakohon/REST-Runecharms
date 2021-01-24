package charm

import "errors"

type Charm struct {
	Id       int    `json:"id"`
	Rune     string `json:"rune"`
	God      string `json:"god"`
	Strength int    `json:"strength"`
}

var (
	ErrNotFound = errors.New("charm not found")
	_           = ErrNotFound
)

// Charm type storage for CRUD
type Storage interface {
	add(charm Charm) (Charm, error)
	get(id int) (Charm, error)
	getAll() ([]Charm, error)
	delete(id int) (Charm, error)
	update(id int, charm Charm) (Charm, error)
}
