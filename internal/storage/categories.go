package storage

import (
	"context"
	"errors"

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
		return category, err
	}

	err = c.db.DB.QueryRow(ctx, sql, args...).Scan(
		&category.ID,
		&category.Name,
		&category.Created,
		&category.Updated,
	)
	if err != nil {
		if errors.Is(pgx.ErrNoRows, err) {
			return category, models.ErrNotFound
		}

		return category, err
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
		return 0, err
	}

	err = c.db.DB.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Delete implements categories.Storager.
func (c *Categories) Delete(ctx context.Context, id int) error {
	sql, args, err := c.psql.Delete("categories").Where("id = ?", id).ToSql()

	if err != nil {
		return err
	}

	result, err := c.db.DB.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return models.ErrNotFound
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
		return err
	}

	result, err := c.db.DB.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return models.ErrNotFound
	}

	return nil
}
