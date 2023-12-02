package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect() {
	var err error

	DATABASE_URL, ok := os.LookupEnv("DATABASE_URL")

	if !ok {
		log.Fatal("DATABASE URL ENV not set")
	}

	pool, err := pgxpool.New(context.Background(), DATABASE_URL)
	if err != nil {
		log.Fatal("Database connection error: ", err)
	}

	DB = pool
}
