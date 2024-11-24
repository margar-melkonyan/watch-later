package app

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"watch-later/internal/storage"

	"github.com/joho/godotenv"
)

func init() {
	ex, _ := os.Getwd()

	err := godotenv.Load(fmt.Sprintf("%s/configs/.env", ex))

	if err != nil {
		log.Panic("The application can't load .env file.")
	}
}

func RunApplication() {
	db, err := storage.NewStorage().Postgres.NewConnection()
	defer func(db *sql.DB) {
		db.Close()
	}(db)

	if err != nil {
		panic(err.Error())
	}

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")),
	}

	slog.Info(fmt.Sprintf(
		"The server running on port: %s",
		os.Getenv("SERVER_PORT"),
	))

	server.ListenAndServe()
}
