package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateMakhdoum(*Makhdoum) error
	DeleteMakhdoum(int) error
	UpdateMakhdoum(*Makhdoum) error
	GetMakhdoumById(int) (*Makhdoum, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres password= dbname=temp sslmode=disable"
	db, err := sql.Open("postgres", connStr)
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
