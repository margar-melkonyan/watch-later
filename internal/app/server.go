package app

import (
	"database/sql"
	"fmt"
	env "github.com/margar-melkonyan/watch-later.git/pkg/env-loader"
	"github.com/margar-melkonyan/watch-later.git/pkg/logs"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/margar-melkonyan/watch-later.git/internal/handler/middleware"
	"github.com/margar-melkonyan/watch-later.git/internal/router"
	"github.com/margar-melkonyan/watch-later.git/internal/storage"
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
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	if err != nil {
		log.Fatal(err.Error())
	}
	slog.Debug("Successfully connected to the db.")

	stack := middleware.Stack(
		middleware.Logging,
	)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")),
		Handler: stack(router.NewRouter()),
	}

	slog.Info(fmt.Sprintf(
		"The server running on port: %s",
		os.Getenv("SERVER_PORT"),
	))
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
