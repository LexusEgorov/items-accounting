package storage

import (
	"context"

	"github.com/Masterminds/squirrel"

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
	errPrefix := "storage.Categories.Get"
	sql, args, err := c.psql.Select("*").From("categories").Where("id = ?", id).ToSql()
	if err != nil {
		return category, c.db.GetError(err, errPrefix)
	}

	err = c.db.DB.QueryRow(ctx, sql, args...).Scan(
		&category.ID,
		&category.Name,
		&category.Created,
		&category.Updated,
	)
	if err != nil {
		return category, c.db.GetError(err, errPrefix)
	}

	return category, nil
}

// Add implements categories.Storager.
func (c *Categories) Add(ctx context.Context, name string) (id int, err error) {
	errPrefix := "storage.Categories.Add"
	sql, args, err := c.psql.Insert("categories").
		Columns("name").
		Values(name).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, c.db.GetError(err, errPrefix)
	}

	err = c.db.DB.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return 0, c.db.GetError(err, errPrefix)
	}

	return id, nil
}

// Delete implements categories.Storager.
func (c *Categories) Delete(ctx context.Context, id int) error {
	errPrefix := "storage.Categories.Delete"
	sql, args, err := c.psql.Delete("categories").Where("id = ?", id).ToSql()

	if err != nil {
		return c.db.GetError(err, errPrefix)
	}

	result, err := c.db.DB.Exec(ctx, sql, args...)
	if err != nil {
		return c.db.GetError(err, errPrefix)
	}

	if result.RowsAffected() == 0 {
		return models.ErrNotFound
	}

	return nil
}

// Set implements categories.Storager.
func (c *Categories) Set(ctx context.Context, id int, name string) error {
	errPrefix := "storage.Categories.Set"
	sql, args, err := c.psql.Update("categories").
		Set("name", name).
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return c.db.GetError(err, errPrefix)
	}

	result, err := c.db.DB.Exec(ctx, sql, args...)
	if err != nil {
		return c.db.GetError(err, errPrefix)
	}

	if result.RowsAffected() == 0 {
		return models.ErrNotUpdated
	}

	return nil
}
