package app

import (
	"database/sql"
	"fmt"
	"github.com/margar-melkonyan/watch-later.git/internal/repository"
	env "github.com/margar-melkonyan/watch-later.git/pkg/env-loader"
	"github.com/margar-melkonyan/watch-later.git/pkg/logs"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/margar-melkonyan/watch-later.git/internal/handler/middleware"
	"github.com/margar-melkonyan/watch-later.git/internal/router"
	"github.com/margar-melkonyan/watch-later.git/internal/storage"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
)

func init() {
	if err := env.MustLoad(); err != nil {
		log.Fatal("The application can't load .env file.")
	}

	if err := logs.Initialize(); err != nil {
		log.Fatal("The applicaiton can't load log level from .env file.")
	}

	m, err := migrate.New(
		os.Getenv("MIGRATION_PATH_URL"),
		fmt.Sprintf(
			"%v://%v:%v@%v:%v/%v?sslmode=%v",
			os.Getenv("DB_DRIVER"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_SSLMODE"),
		))

	if err != nil {
		log.Fatal(fmt.Sprintf("The applicaiton can't initialize migration instance. error: {%v}", err.Error()))
	}

	_ = m.Up()
}

func RunApplication() {
	slog.Debug("Connection to the db...")
	db, err := storage.NewStorage().Postgres.NewConnection()
	defer func(db *sql.DB) {
		db.Close()
	}(db)

	userRepository := repository.NewUserRepository(db)
	userRepository.Create(&repository.User{
		Nickname:   "12332",
		Firstname:  "test",
		Lastname:   "234",
		Patronymic: "123",
		Email:      "23",
		Password:   "3434",
	})
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
