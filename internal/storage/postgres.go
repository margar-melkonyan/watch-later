package storage

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Postgres struct{}

func (connection Postgres) NewConnection() (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := sql.Open(os.Getenv("DB_DRIVER"), dsn)

	if err != nil {
		return nil, fmt.Errorf(
			"The applicaiton cannot open connection with %s",
			os.Getenv("DB_DRIVER"),
		)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("The application cannot connect to the database.")
	}

	return db, nil
}
