package inmem

import (
	"github.com/yuriimakohon/RunecharmsCRUD/api/models"
	"testing"
)

var s = New()

func TestStorage_Add(t *testing.T) {
	c1 := models.Charm{
		Id:    0,
		Rune:  "Mannaz",
		God:   "Tyr",
		Power: 100,
	}
	c2 := models.Charm{
		Id:    5,
		Rune:  "Sunszu",
		God:   "Odin",
		Power: 200,
	}

	r1, err := s.Add(c1)
	if err != nil || c1 != r1 {
		t.Fatalf("ERR: %s | %v mismatchs %v\n", err, c1, r1)
	}

	r2, err := s.Add(c2)
	c2.Id = 1
	if err != nil || c2 != r2 {
		t.Fatalf("ERR: %s | %v mismatchs %v\n", err, c2, r2)
	}
}
