package charm

type IMCharms struct {
	storage []Charm
	lastId  int
}

func (s *IMCharms) add(c Charm) (Charm, error) {
	c.Id = s.lastId
	s.storage = append(s.storage, c)
	s.lastId++
	return c, nil
}

func (s *IMCharms) getAll() ([]Charm, error) {
	cp := make([]Charm, len(s.storage))
	copy(cp, s.storage)
	return cp, nil
}

func (s *IMCharms) get(id int) (Charm, error) {
	for _, c := range s.storage {
		if c.Id == id {
			return c, nil
		}
	}
	return Charm{}, ErrNotFound
}

func (s *IMCharms) delete(id int) (Charm, error) {
	for idx, c := range s.storage {
		if c.Id == id {
			s.storage = append(s.storage[:idx], s.storage[idx+1:]...)
			return c, nil
		}
	}
	return Charm{}, ErrNotFound
}

func (s *IMCharms) update(id int, u Charm) (Charm, error) {
	for idx, c := range s.storage {
		if c.Id == id {
			u.Id = id
			s.storage[idx] = u
			return s.storage[idx], nil
		}
	}
	return Charm{}, ErrNotFound
}
