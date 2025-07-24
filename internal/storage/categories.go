package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"github.com/LexusEgorov/items-accounting/internal/models"
)

type Categories struct {
	db   *DB
	psql squirrel.StatementBuilderType
}

func NewCategories(db *DB) (*Categories, error) {
	return &Categories{
		db:   db,
		psql: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}, nil
}

// Get implements categories.Storager.
func (c *Categories) Get(ctx context.Context, id int) (category models.Category, err error) {
	sql, args, err := c.psql.Select("*").From("categories").Where("id = ?", id).ToSql()
	if err != nil {
		return category, fmt.Errorf("Storage.Categories.Get: %v", err)
	}

	err = c.db.DB.QueryRow(ctx, sql, args...).Scan(
		&category.ID,
		&category.Name,
		&category.Created,
		&category.Updated,
	)
	if err != nil {
		if errors.Is(pgx.ErrNoRows, err) {
			return category, fmt.Errorf("Storage.Categories.Get: %v", models.ErrNotFound)
		}

		return category, fmt.Errorf("Storage.Categories.Get: %v", err)
	}

	return category, nil
}

// Add implements categories.Storager.
func (c *Categories) Add(ctx context.Context, name string) (id int, err error) {
	sql, args, err := c.psql.Insert("categories").
		Columns("name").
		Values(name).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("Storage.Categories.Add: %v", err)
	}

	err = c.db.DB.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("Storage.Categories.Add: %v", err)
	}

	return id, nil
}

// Delete implements categories.Storager.
func (c *Categories) Delete(ctx context.Context, id int) error {
	sql, args, err := c.psql.Delete("categories").Where("id = ?", id).ToSql()

	if err != nil {
		return fmt.Errorf("Storage.Categories.Delete: %v", err)
	}

	result, err := c.db.DB.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("Storage.Categories.Delete: %v", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("Storage.Categories.Delete: %v", models.ErrNotFound)
	}

	return nil
}

// Set implements categories.Storager.
func (c *Categories) Set(ctx context.Context, id int, name string) error {
	sql, args, err := c.psql.Update("categories").
		Set("name", name).
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return fmt.Errorf("Storage.Categories.Set: %v", err)
	}

	result, err := c.db.DB.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("Storage.Categories.Set: %v", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("Storage.Categories.Get: %v", models.ErrNotFound)
	}

	return nil
}
