package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/LexusEgorov/items-accounting/internal/config"
)

type Storager interface {
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Ping(ctx context.Context) error
	Close()
}

type DB struct {
	DB Storager
}

func NewDB(cfg config.DBConfig) (*DB, error) {
	db := &DB{}

	connStr := config.GetConnStr(cfg.User, cfg.Password, cfg.Name)
	err := db.connect(connStr)
	if err != nil {
		return nil, fmt.Errorf("newDb: %v", err)
	}

	return db, nil
}

func (d *DB) Close() {
	d.DB.Close()
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
