package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"github.com/LexusEgorov/items-accounting/internal/models"
)

type Products struct {
	db   *DB
	psql squirrel.StatementBuilderType
}

func NewProducts(db *DB) (*Products, error) {
	return &Products{
		db:   db,
		psql: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}, nil
}

// Add implements products.Storager.
func (p *Products) Add(ctx context.Context, product models.ProductDTO) (id int, err error) {
	sql, args, err := p.psql.Insert("products").
		Columns("category_id", "name", "price", "count").
		Values(product.CatID, product.Name, product.Price, product.Count).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("Storage.Products.Add: %v", err)
	}

	err = p.db.DB.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("Storage.Products.Add: %v", err)
	}

	return
}

// Delete implements products.Storager.
func (p *Products) Delete(ctx context.Context, id int) error {
	sql, args, err := p.psql.Delete("products").Where("id = ?", id).ToSql()
	if err != nil {
		return fmt.Errorf("Storage.Products.Delete: %v", err)
	}

	result, err := p.db.DB.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("Storage.Products.Delete: %v", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("Storage.Products.Delete: %v", err)
	}

	return nil
}

// Get implements products.Storager.
func (p *Products) Get(ctx context.Context, id int) (product models.Product, err error) {
	sql, args, err := p.psql.Select("*").From("products").Where("id = ?", id).ToSql()
	if err != nil {
		return product, fmt.Errorf("Storage.Products.Get: %v", err)
	}

	err = p.db.DB.QueryRow(ctx, sql, args...).Scan(
		&product.ID,
		&product.CatID,
		&product.Name,
		&product.Price,
		&product.Count,
		&product.Created,
		&product.Updated,
	)
	if err != nil {
		if errors.Is(pgx.ErrNoRows, err) {
			return product, fmt.Errorf("Storage.Products.Get: %v", models.ErrNotFound)
		}

		return product, fmt.Errorf("Storage.Products.Get: %v", err)
	}

	return product, nil
}

// Set implements products.Storager.
func (p *Products) Set(ctx context.Context, product models.ProductDTO) error {
	sql, args, err := p.psql.Update("products").
		Set("category_id", product.CatID).
		Set("name", product.Name).
		Set("price", product.Price).
		Set("count", product.Count).
		Where("id = ?", product.ID).
		ToSql()
	if err != nil {
		return fmt.Errorf("Storage.Products.Set: %v", err)
	}

	result, err := p.db.DB.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("Storage.Products.Get: %v", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("Storage.Products.Get: %v", err)
	}

	return nil
}
