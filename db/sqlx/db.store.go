package db

import (
	"github.com/jmoiron/sqlx"
)

type Store interface {
	Querier
}

type SQLStore struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) Store {
	return &SQLStore{
		db: db,
	}
}
