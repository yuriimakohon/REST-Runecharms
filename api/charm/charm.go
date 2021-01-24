package charm

type Charm struct {
	Id       int    `json:"id"`
	Rune     string `json:"rune"`
	God      string `json:"god"`
	Strength int    `json:"strength"`
}

// Charm type storage for CRUD
type Storage interface {
	add(charm Charm) *Charm
	get(id int) *Charm
	getAll() []Charm
	delete(id int)
	update(id int, charm Charm) *Charm
}
