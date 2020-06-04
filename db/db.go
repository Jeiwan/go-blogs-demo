package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	db *sqlx.DB
}

func New(datasource *Datasource) (*DB, error) {
	db, err := sqlx.Connect("postgres", datasource.String())
	if err != nil {
		return nil, err
	}

	return &DB{db: db}, nil
}
