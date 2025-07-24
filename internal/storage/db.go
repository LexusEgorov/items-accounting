package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/LexusEgorov/items-accounting/internal/config"
)

type DB struct {
	DB *pgxpool.Pool
}

func NewDB(cfg config.DBConfig) (*DB, error) {
	db := &DB{}

	connStr := config.GetConnStr(cfg.User, cfg.Password, cfg.Name)
	err := db.connect(connStr)
	if err != nil {
		return nil, fmt.Errorf("Storage.db.NewDB: %v", err)
	}

	return db, nil
}

func (d *DB) Close() {
	d.DB.Close()
}

func (d *DB) connect(connStr string) error {
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return fmt.Errorf("Storage.db.connect: %v", err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		return fmt.Errorf("Storage.db.connect: %v", err)
	}

	d.DB = pool
	return nil
}
