package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/LexusEgorov/items-accounting/internal/models"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgconn"
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
		return category, fmt.Errorf("categoriesDb.get: %v", err)
	}

	err = c.db.DB.QueryRow(context.TODO(), sql, args...).Scan(&category)

	if err != nil {
		return category, fmt.Errorf("categoriesDb.get: %v", err)
	}

	return category, nil
}

// Add implements categories.Storager.
func (c *Categories) Add(name string) (id int, err error) {
	sql, args, err := c.psql.Insert("categories").
		Columns("name").
		Values(name).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return 0, fmt.Errorf("add category: %v", err)
	}

	err = c.db.DB.QueryRow(context.TODO(), sql, args...).Scan(&id)

	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == models.ErrUniqueCode {
			return 0, models.ErrUnique
		}

		return 0, fmt.Errorf("add category: %v", err)
	}

	return id, nil
}

// Delete implements categories.Storager.
func (c *Categories) Delete(id int) error {
	sql, args, err := c.psql.Delete("categories").Where("id = ?", id).ToSql()

	if err != nil {
		return fmt.Errorf("delete category: %v", err)
	}

	result, err := c.db.DB.Exec(context.TODO(), sql, args...)
	if err != nil {
		return fmt.Errorf("delete category: %v", err)
	}

	if result.RowsAffected() == 0 {
		return models.ErrNotFound
	}

	return nil
}

// Set implements categories.Storager.
func (c *Categories) Set(id int, name string) error {
	sql, args, err := c.psql.Update("categories").
		Set("name", name).
		Where("id = ?", id).
		ToSql()

	if err != nil {
		return fmt.Errorf("set category: %v", err)
	}

	result, err := c.db.DB.Exec(context.TODO(), sql, args...)

	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == models.ErrUniqueCode {
			return models.ErrUnique
		}
	}

	if result.RowsAffected() == 0 {
		return models.ErrNotUpdated
	}

	return nil
}
