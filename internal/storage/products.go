package storage

import (
	"github.com/LexusEgorov/items-accounting/internal/models"
	"github.com/Masterminds/squirrel"
)

type Products struct {
	db   *DB
	psql squirrel.StatementBuilderType
}

// TODO: add logic for methods
func NewProducts(connStr string) (*Products, error) {
	db, err := NewDB(connStr)

	if err != nil {
		return nil, err
	}

	return &Products{
		db:   db,
		psql: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}, nil
}

// Add implements products.Storager.
func (p *Products) Add(string) (models.Product, error) {
	panic("unimplemented")
}

// Delete implements products.Storager.
func (p *Products) Delete(int) error {
	panic("unimplemented")
}

// Get implements products.Storager.
func (p *Products) Get(int) (models.Product, error) {
	panic("unimplemented")
}

// Set implements products.Storager.
func (p *Products) Set(models.Product) (models.Product, error) {
	panic("unimplemented")
}
