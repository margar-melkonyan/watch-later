package app

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/margar-melkonyan/watch-later.git/internal/handler/middleware"
	"github.com/margar-melkonyan/watch-later.git/internal/router"
	"github.com/margar-melkonyan/watch-later.git/internal/storage"
	env "github.com/margar-melkonyan/watch-later.git/pkg/env-loader"
	"github.com/margar-melkonyan/watch-later.git/pkg/logs"
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
		middleware.AuthMiddleware,
		middleware.Logging,
		middleware.SetLocale,
	)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")),
		Handler: stack(router.NewRouter(db)),
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	go func() {
		slog.Info(fmt.Sprintf(
			"The server running on port: %s",
			os.Getenv("SERVER_PORT"),
		))
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err.Error())
		}
	}()

	<-ctx.Done()
}
