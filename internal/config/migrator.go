package config

import (
	"flag"

	"github.com/LexusEgorov/items-accounting/internal/models"
)

type MigratorConfig struct {
	DBConfig
	MigrationsPath string
}

func NewMigratorConfig() (*MigratorConfig, error) {
	var (
		migrationsPath string
		user           string
		password       string
		name           string
	)

	flag.StringVar(&migrationsPath, "m", "", "path to migrations")
	flag.StringVar(&password, "p", "", "db password")
	flag.StringVar(&name, "n", "", "db name")
	flag.StringVar(&user, "u", "", "db username")
	flag.Parse()

	if migrationsPath == "" {
		return nil, models.ErrMigrationsNotProvided
	}

	if user == "" {
		return nil, models.ErrBadUserName
	}

	if password == "" {
		return nil, models.ErrBadPassword
	}

	if name == "" {
		return nil, models.ErrBadDBName
	}

	return &MigratorConfig{
		MigrationsPath: migrationsPath,
		DBConfig: DBConfig{
			User:     user,
			Password: password,
			Name:     name,
		},
	}, nil
}
