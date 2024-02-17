package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore(user, name, pass string) (*PostgresStore, error) {
	connectionString := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", user, name, pass)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.createTodoTable()
}

func (s *PostgresStore) createTodoTable() error {
	query := `CREATE TABLE IF NOT EXISTS todo (
	id serial primary key,
	title varchar(50) not null, 
	status boolean not null default false,
	created_at timestamp not null
	)`

	_, err := s.db.Exec(query)
	return err
}
