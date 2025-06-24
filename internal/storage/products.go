package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/LexusEgorov/items-accounting/internal/models"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgconn"
)

type Products struct {
	db   *DB
	psql squirrel.StatementBuilderType
}

// TODO: add logic for methods
func NewProducts(db *DB) (*Products, error) {
	return &Products{
		db:   db,
		psql: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}, nil
}

// Add implements products.Storager.
func (p *Products) Add(product models.ProductDTO) (id int, err error) {
	sql, args, err := p.psql.Insert("products").
		Columns("category_id", "name", "price", "count").
		Values(product.CatID, product.Name, product.Price, product.Count).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return 0, fmt.Errorf("product add: %v", err)
	}

	err = p.db.DB.QueryRow(context.TODO(), sql, args...).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("product add: %v", err)
	}

	return id, err
}

// Delete implements products.Storager.
func (p *Products) Delete(id int) error {
	sql, args, err := p.psql.Delete("products").Where("id = ?", id).ToSql()

	if err != nil {
		return fmt.Errorf("delete product: %v", err)
	}

	result, err := p.db.DB.Exec(context.TODO(), sql, args...)
	if err != nil {
		return fmt.Errorf("delete product: %v", err)
	}

	if result.RowsAffected() == 0 {
		return models.ErrNotFound
	}

	return nil
}

// Get implements products.Storager.
func (p *Products) Get(id int) (product models.Product, err error) {
	sql, args, err := p.psql.Select("*").From("products").Where("id = ?", id).ToSql()

	if err != nil {
		return product, fmt.Errorf("get product: %v", err)
	}

	err = p.db.DB.QueryRow(context.TODO(), sql, args...).Scan(&product)

	if err != nil {
		return product, fmt.Errorf("get product: %v", err)
	}

	return product, nil
}

// Set implements products.Storager.
func (p *Products) Set(product models.ProductDTO) error {
	sql, args, err := p.psql.Update("categories").
		Set("category_id", product.CatID).
		Set("name", product.Name).
		Set("price", product.Price).
		Set("count", product.Count).
		Where("id = ?", product.ID).
		ToSql()

	if err != nil {
		return fmt.Errorf("set product: %v", err)
	}

	result, err := p.db.DB.Exec(context.TODO(), sql, args...)

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
