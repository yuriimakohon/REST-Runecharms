package grpc

import (
	m "github.com/yuriimakohon/RunecharmsCRUD/api/models"
	a "github.com/yuriimakohon/RunecharmsCRUD/pkg/api"
)

func CharmGrpcToModel(c *a.Charm) m.Charm {
	if c != nil {
		return m.Charm{
			Id:    c.GetId(),
			Rune:  c.GetRune(),
			God:   c.GetGod(),
			Power: c.GetPower(),
		}
	}
	return m.Charm{}
}

func CharmModelToGrpc(c m.Charm) *a.Charm {
	return &a.Charm{
		Id:    c.Id,
		Rune:  c.Rune,
		God:   c.God,
		Power: c.Power,
	}
}
