package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var err error
var db *pgxpool.Pool

func Connect() error {

	DATABASE_URL, ok := os.LookupEnv("DATABASE_URL")

	if !ok {
		panic("DATABASE URL ENV not set")
	}

	Migrate()

	pool, err := pgxpool.New(context.Background(), DATABASE_URL)
	db = pool

	if err != nil {
		log.Fatal("Database connection error: ", err)
	}

	return nil
}
