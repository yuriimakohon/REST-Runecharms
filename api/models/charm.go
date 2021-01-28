package models

type Charm struct {
	Id    int    `json:"id"`
	Rune  string `json:"rune"`
	God   string `json:"god"`
	Power int    `json:"power"`
}
