package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/LexusEgorov/items-accounting/internal/config"
	"github.com/LexusEgorov/items-accounting/internal/models"
	"github.com/LexusEgorov/items-accounting/internal/utils"
)

type DB struct {
	DB *pgxpool.Pool
}

func NewDB(config config.DBConfig) (*DB, error) {
	db := &DB{}

	connStr := utils.GetConnStr(config.User, config.Password, config.Name)
	err := db.connect(connStr)
	if err != nil {
		return nil, fmt.Errorf("newDb: %v", err)
	}

	return db, nil
}

func (d *DB) Close() {
	d.DB.Close()
}

func (d *DB) GetError(err error, prefix string) error {
	var pgErr *pgconn.PgError

	if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
		err = models.ErrUnique
	}

	return fmt.Errorf("%s: %v", prefix, err)
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
