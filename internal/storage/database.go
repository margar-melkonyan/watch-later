package storage

import (
	"database/sql"
)

type DBConnector interface {
	NewConnection() (*sql.DB, error)
}

type Database struct {
	Postgres DBConnector
}

func NewStorage() *Database {
	return &Database{
		Postgres: &Postgres{},
	}
}
