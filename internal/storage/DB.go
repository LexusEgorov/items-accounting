package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	DB *pgxpool.Pool
}

func NewDB(connStr string) (*DB, error) {
	db := &DB{}

	err := db.connect(connStr)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func (d *DB) connect(connStr string) error {
	pool, err := pgxpool.New(context.Background(), connStr)

	if err != nil {
		return fmt.Errorf("db.connect: %v", err)
	}

	err = pool.Ping(context.Background())

	if err != nil {
		return fmt.Errorf("db.connect pool ping: %v", err)
	}

	d.DB = pool
	return nil
}

func (d *DB) Close() {
	d.DB.Close()
}
