package app

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
)

func RunMigrations() {
	slog.Info("Running migrations...")
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
	slog.Info("Migration applied successfully")
}
