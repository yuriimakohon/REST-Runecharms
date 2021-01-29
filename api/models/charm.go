package models

type Charm struct {
	Id    int32  `json:"id"`
	Rune  string `json:"rune"`
	God   string `json:"god"`
	Power int32  `json:"power"`
}
