package storage

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type DB struct {
	DB *pgx.Conn
}

var db *DB

func NewDB(connStr string) (*DB, error) {
	if db != nil {
		return db, nil
	}

	db = &DB{}

	err := db.connect(connStr)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func (d *DB) connect(connStr string) error {
	conn, err := pgx.Connect(context.Background(), connStr)

	if err != nil {
		return err
	}

	err = conn.Ping(context.Background())

	if err != nil {
		return err
	}

	d.DB = conn

	return nil
}

func (d *DB) Close(ctx context.Context) error {
	return d.DB.Close(ctx)
}
