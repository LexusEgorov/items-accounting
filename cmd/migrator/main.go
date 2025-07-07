package main

import (
	"log"

	"github.com/LexusEgorov/items-accounting/internal/config"
	"github.com/LexusEgorov/items-accounting/internal/utils"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func main() {
	cfg, err := config.NewMigratorConfig()
	if err != nil {
		log.Fatalf("migrator: %v", err)
	}

	connStr := utils.GetConnStr(cfg.User, cfg.Password, cfg.Name)
	connConfig, err := pgx.ParseConfig(connStr)
	if err != nil {
		log.Fatalf("migrator: failed to parse conn config: %v", err)
	}

	db := stdlib.OpenDB(*connConfig)
	defer db.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("migrator: goose error: %v", err)
	}

	if err := goose.Up(db, cfg.MigrationsPath); err != nil {
		log.Fatalf("migrator: goose error: %v", err)
	}
}
