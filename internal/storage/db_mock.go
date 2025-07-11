package storage

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type mockDB struct{}
type mockRow struct{}

func (m mockRow) Scan(dest ...any) error {
	return nil
}

func (m mockDB) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return mockRow{}
}

func (m mockDB) Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
	return pgconn.NewCommandTag("mock"), nil
}

func (m mockDB) Ping(ctx context.Context) error {
	return nil
}

func (m mockDB) Close() {}
