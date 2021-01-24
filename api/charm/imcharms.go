package charm

type IMCharms struct {
	storage []Charm
	lastId  int
}

func (s *IMCharms) add(c Charm) *Charm {
	c.Id = s.lastId
	s.storage = append(s.storage, c)
	s.lastId++
	return &c
}

func (s *IMCharms) getAll() []Charm {
	return s.storage
}

func (s *IMCharms) get(id int) *Charm {
	for _, c := range s.storage {
		if c.Id == id {
			return &c
		}
	}
	return nil
}

func (s *IMCharms) delete(id int) {
	for idx, c := range s.storage {
		if c.Id == id {
			s.storage = append(s.storage[:idx], s.storage[idx+1:]...)
		}
	}
}

func (s *IMCharms) update(id int, u Charm) *Charm {
	for idx, c := range s.storage {
		if c.Id == id {
			u.Id = id
			s.storage[idx] = u
			return &s.storage[idx]
		}
	}
	return nil
}
