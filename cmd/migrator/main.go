package main

import (
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

// TODO: add config for connection
func main() {
	connStr := "postgres://root:root@db:5432/accounting-db?sslmode=disable"
	connConfig, err := pgx.ParseConfig(connStr)

	if err != nil {
		log.Fatalf("migrator: failed to parse conn config: %v", err)
	}

	db := stdlib.OpenDB(*connConfig)
	defer db.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("migrator: goose error: %v", err)
	}

	if err := goose.Up(db, "/app/migrations"); err != nil {
		log.Fatalf("migrator: goose error: %v", err)
	}
}
