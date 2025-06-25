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

// TODO: add logic for all methods
func NewCategories(db *DB) (*Categories, error) {
	return &Categories{
		db:   db,
		psql: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}, nil
}

// Get implements categories.Storager.
func (c *Categories) Get(id int) (category models.Category, err error) {
	errPrefix := "storage.Categories.Get"
	sql, args, err := c.psql.Select("*").From("categories").Where("id = ?", id).ToSql()

	if err != nil {
		return category, c.db.GetError(err, errPrefix)
	}

	err = c.db.DB.QueryRow(context.TODO(), sql, args...).Scan(&category)

	if err != nil {
		return category, c.db.GetError(err, errPrefix)
	}

	return category, nil
}

// Add implements categories.Storager.
func (c *Categories) Add(name string) (id int, err error) {
	errPrefix := "storage.Categories.Add"
	sql, args, err := c.psql.Insert("categories").
		Columns("name").
		Values(name).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return 0, c.db.GetError(err, errPrefix)
	}

	err = c.db.DB.QueryRow(context.TODO(), sql, args...).Scan(&id)

	if err != nil {
		return 0, c.db.GetError(err, errPrefix)
	}

	return id, nil
}

// Delete implements categories.Storager.
func (c *Categories) Delete(id int) error {
	errPrefix := "storage.Categories.Delete"
	sql, args, err := c.psql.Delete("categories").Where("id = ?", id).ToSql()

	if err != nil {
		return c.db.GetError(err, errPrefix)
	}

	result, err := c.db.DB.Exec(context.TODO(), sql, args...)
	if err != nil {
		return c.db.GetError(err, errPrefix)
	}

	if result.RowsAffected() == 0 {
		return models.ErrNotFound
	}

	return nil
}

// Set implements categories.Storager.
func (c *Categories) Set(id int, name string) error {
	errPrefix := "storage.Categories.Set"
	sql, args, err := c.psql.Update("categories").
		Set("name", name).
		Where("id = ?", id).
		ToSql()

	if err != nil {
		return c.db.GetError(err, errPrefix)
	}

	result, err := c.db.DB.Exec(context.TODO(), sql, args...)

	if err != nil {
		return c.db.GetError(err, errPrefix)
	}

	if result.RowsAffected() == 0 {
		return models.ErrNotUpdated
	}

	return nil
}
