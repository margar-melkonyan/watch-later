package app

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"watch-later/internal/storage"
	env "watch-later/pkg/env-loader"
	"watch-later/pkg/logs"
)

func init() {
	if err := env.MustLoad(); err != nil {
		log.Fatal("The application can't load .env file.")
	}

	if err := logs.Initialize(); err != nil {
		log.Fatal("The applicaiton can't load log level from .env file.")
	}
}

func RunApplication() {
	slog.Debug("Connection to the db...")
	db, err := storage.NewStorage().Postgres.NewConnection()
	defer func(db *sql.DB) {
		db.Close()
	}(db)

	if err != nil {
		log.Fatal(err.Error())
	}
	slog.Debug("Successfully connected to the db.")

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")),
	}

	slog.Info(fmt.Sprintf(
		"The server running on port: %s",
		os.Getenv("SERVER_PORT"),
	))

	server.ListenAndServe()
}
