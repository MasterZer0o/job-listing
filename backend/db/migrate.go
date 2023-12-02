package db

import (
	"log/slog"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(forceVer *int) {
	m, err := migrate.New(
		"file://db/migrations",
		os.Getenv("DATABASE_URL"))

	if forceVer != nil {
		m.Force(*forceVer)
	}

	if err != nil {
		slog.Error(err.Error())
		return
	}

	if err := m.Up(); err != nil {
		slog.Error("Migration error: ", "", err.Error())
		return
	}

	slog.Info("Migration successful.")
}
