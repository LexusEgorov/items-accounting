package config

import (
	"flag"

	"github.com/LexusEgorov/items-accounting/internal/models"
)

type MigratorConfig struct {
	MigrationsPath string `yaml:"migrationsPath"`
}

func NewMigratorConfig() (*MigratorConfig, error) {
	var migrationsPath string

	flag.StringVar(&migrationsPath, "m", "", "path to migrations")
	flag.Parse()

	if migrationsPath == "" {
		return nil, models.ErrMigrationsNotProvided
	}

	return &MigratorConfig{
		MigrationsPath: migrationsPath,
	}, nil
}
