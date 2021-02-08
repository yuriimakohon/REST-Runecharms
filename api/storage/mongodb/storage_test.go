package mongodb

import (
	"fmt"
	"github.com/yuriimakohon/RunecharmsCRUD/api/models"
	"testing"
)

var s = New()

func TestStartClean(t *testing.T) {
	s.coll.Drop(s.ctx)
}

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

func TestStorage_Get(t *testing.T) {
	c1 := models.Charm{
		Id:    55,
		Rune:  "Kilao",
		God:   "Loki",
		Power: 140,
	}

	c1, err := s.Add(c1)
	if err != nil {
		t.Fatalf("ERR: %s | %v \n", err, c1)
		return
	}

	r1, err := s.Get(c1.Id)
	if err != nil {
		t.Fatalf("ERR: %s | %v \n", err, c1)
		return
	}

	if r1 != c1 {
		t.Fatalf("ERR: %s | %v mismatchs %v\n", err, c1, r1)
	}
}

func TestStorage_GetAll(t *testing.T) {
	sli, err := s.GetAll()
	if err != nil || sli == nil {
		t.Fatalf("ERR: %s | %v", err, sli)
	}
}

func TestStorage_Delete(t *testing.T) {
	c1 := models.Charm{
		Id:    45,
		Rune:  "Mannaz",
		God:   "Freya",
		Power: 777,
	}

	c1, err := s.Add(c1)
	if err != nil {
		t.Fatalf("ERR: %s | %v \n", err, c1)
		return
	}

	r1, err := s.Delete(c1.Id)
	if err != nil || r1 != c1 {
		t.Fatalf("ERR: %s | %v mismatchs %v\n", err, c1, r1)
		return
	}

	r1, err = s.Get(c1.Id)
	zeroCharm := models.Charm{}
	if err == nil && r1 != zeroCharm {
		t.Fatalf("ERR: %s | wasn't deleted %v\n", err, r1)
	}
}

func TestStorage_Update(t *testing.T) {
	c1 := models.Charm{
		Id:    -5,
		Rune:  "TestRune",
		God:   "TestGod",
		Power: -56,
	}

	r1, err := s.Update(0, c1)
	c1.Id = 0
	if err != nil {
		t.Fatalf("ERR: %s | %v\n", err, r1)
		return
	}

	r2, err := s.Get(0)
	if err != nil {
		t.Fatalf("ERR: %s | %v\n", err, r1)
		return
	}

	if c1 != r1 || c1 != r2 {
		t.Fatalf("mismatchs %v : %v : %v\n", c1, r1, r2)
	}
}

func TestStorage_Len(t *testing.T) {
	l, err := s.Len()
	if err != nil {
		t.Fatalf("ERR : %s\n", err)
		return
	}

	if l > 0 {
		sli, err := s.GetAll()
		if err != nil {
			t.Fatalf("ERR : %s\n", err)
			return
		}
		for _, c := range sli {
			if _, err = s.Delete(c.Id); err != nil {
				t.Fatalf("ERR : %s | %v\n", err, c)
			}
		}

		l, err = s.Len()
		if err != nil {
			t.Fatalf("ERR : %s\n", err)
			return
		}
		if l != 0 {
			t.Fatalf("Len != 0 : %v", l)
			return
		}
	} else if l == 0 {
		c1 := models.Charm{
			Id:    787,
			Rune:  "Mannaz",
			God:   "Freya",
			Power: -5,
		}

		count := 5000
		for i := 0; i < count; i++ {
			if _, err = s.Add(c1); err != nil {
				t.Fatalf("ERR : %s\n", err)
				return
			}
		}

		l, err = s.Len()
		fmt.Println(l)
		if err != nil {
			t.Fatalf("ERR : %s\n", err)
			return
		}
		if l != count {
			t.Fatalf("must be %v: %v\n", count, l)
		}

	} else {
		t.Fatalf("Invalid Len result: %v", l)
	}
}

func TestFinishClean(t *testing.T) {
	s.coll.Drop(s.ctx)
}
