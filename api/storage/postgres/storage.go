package postgres

import (
	"context"
	"github.com/jackc/pgx/v4"
	m "github.com/yuriimakohon/RunecharmsCRUD/api/models"
	"github.com/yuriimakohon/RunecharmsCRUD/api/storage"
	"log"
)

type Storage struct {
	ctx  context.Context
	conn *pgx.Conn
}

const createSql = `
CREATE TABLE IF NOT EXISTS charms(
id SERIAL,
rune varchar(255),
god varchar(255),
power int);`

func New() *Storage {
	st := &Storage{ctx: context.Background()}
	conn, err := pgx.Connect(st.ctx, "postgresql://postgres:secret@localhost:5432/postgres")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		return nil
	}

	if err = conn.Ping(st.ctx); err != nil {
		log.Fatal(err)
		return nil
	}
	st.conn = conn

	rows, _ := st.conn.Query(st.ctx, createSql)
	rows.Close()

	return st
}

func (s *Storage) Add(charm m.Charm) (m.Charm, error) {
	query := `INSERT INTO charms (rune, god, power) VALUES ($1, $2, $3) RETURNING id`

	rows, err := s.conn.Query(s.ctx, query, charm.Rune, charm.God, charm.Power)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
		return m.Charm{}, err
	}

	rows.Next()
	if err = rows.Scan(&charm.Id); err != nil {
		log.Fatal(err)
		return m.Charm{}, err
	}

	return charm, nil
}

func (s *Storage) Get(id int32) (m.Charm, error) {
	query := `SELECT * FROM charms WHERE id=$1`
	charm := m.Charm{}

	rows, err := s.conn.Query(s.ctx, query, id)
	defer rows.Close()
	if err != nil {
		return m.Charm{}, storage.ErrNotFound
	}

	if !rows.Next() {
		return m.Charm{}, storage.ErrNotFound
	}

	if err = rows.Scan(&charm.Id, &charm.Rune, &charm.God, &charm.Power); err != nil {
		log.Fatal(err)
		return m.Charm{}, err
	}
	return charm, nil
}

func (s *Storage) GetAll() ([]m.Charm, error) {
	charms := make([]m.Charm, 0, 1)

	rows, err := s.conn.Query(s.ctx, "SELECT * FROM charms")
	if err != nil {
		return []m.Charm{}, err
	}
	defer rows.Close()

	charm := m.Charm{}
	for rows.Next() {
		if err = rows.Scan(&charm.Id, &charm.Rune, &charm.God, &charm.Power); err != nil {
			log.Fatal(err)
		} else {
			charms = append(charms, charm)
		}
	}
	return charms, nil
}

func (s *Storage) Delete(id int32) (m.Charm, error) {
	query := `DELETE FROM charms WHERE id=$1`

	charm, err := s.Get(id)
	if err != nil {
		return m.Charm{}, err
	}

	rows, err := s.conn.Query(s.ctx, query, id)
	if err != nil {
		return m.Charm{}, storage.ErrNotFound
	}
	rows.Close()

	return charm, nil
}

func (s *Storage) Update(id int32, charm m.Charm) (m.Charm, error) {
	query := `UPDATE charms SET rune=$2, god=$3, power=$4 WHERE id=$1`

	if _, err := s.Get(id); err != nil {
		return m.Charm{}, err
	}

	rows, err := s.conn.Query(s.ctx, query, id, charm.Rune, charm.God, charm.Power)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
		return m.Charm{}, err
	}

	charm.Id = id
	return charm, nil
}

func (s *Storage) Len() (int, error) {
	length := 0
	rows, err := s.conn.Query(s.ctx, "SELECT COUNT(*) FROM charms")
	if err != nil {
		return 0, err
	}

	rows.Next()
	defer rows.Close()
	if err = rows.Scan(&length); err != nil {
		return 0, err
	}
	return length, nil
}
