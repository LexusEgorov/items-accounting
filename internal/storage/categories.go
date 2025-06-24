package storage

import (
	"context"

	"github.com/LexusEgorov/items-accounting/internal/models"
	"github.com/Masterminds/squirrel"
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
func (c *Categories) Get(ID int) (models.Category, error) {
	category := models.Category{}
	sql, args, err := c.psql.Select("*").From("categories").Where("id = ?", ID).ToSql()

	if err != nil {
		return category, err
	}

	err = c.db.DB.QueryRow(context.Background(), sql, args...).Scan(&category)

	if err != nil {
		return category, err
	}

	return category, nil
}

// Add implements categories.Storager.
func (c *Categories) Add(string) (models.Category, error) {
	panic("unimplemented")
}

// Delete implements categories.Storager.
func (c *Categories) Delete(int) error {
	panic("unimplemented")
}

// Set implements categories.Storager.
func (c *Categories) Set(int, string) (models.Category, error) {
	panic("unimplemented")
}
