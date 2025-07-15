package products

import (
	"context"
	"fmt"

	"github.com/LexusEgorov/items-accounting/internal/models"
)

type ProductRepository interface {
	Add(ctx context.Context, product models.ProductDTO) (id int, err error)
	Get(ctx context.Context, id int) (product models.Product, err error)
	Set(ctx context.Context, product models.ProductDTO) error
	Delete(ctx context.Context, id int) error
}

type Products struct {
	storage ProductRepository
}

func New(storage ProductRepository) *Products {
	return &Products{
		storage: storage,
	}
}

func (c Products) Add(ctx context.Context, product models.ProductDTO) (models.ProductDTO, error) {
	//TODO: move to validateFn
	if product.CatID == 0 {
		return models.ProductDTO{}, fmt.Errorf("Products.Add: %v", models.NewEmptyErr("categoryId"))
	}

	//TODO: move to validateFn
	if product.Name == "" {
		return models.ProductDTO{}, fmt.Errorf("Products.Add: %v", models.NewEmptyErr("name"))
	}

	id, err := c.storage.Add(ctx, product)
	if err != nil {
		return models.ProductDTO{}, fmt.Errorf("Products.Add: %v", err)
	}

	product.ID = id
	return product, nil
}

func (c Products) Set(ctx context.Context, product models.ProductDTO) (models.ProductDTO, error) {
	//TODO: move to validateFn
	if product.CatID <= 0 {
		return models.ProductDTO{}, fmt.Errorf("Products.Set: %v", models.NewEmptyErr("categoryId"))
	}

	//TODO: move to validateFn
	if product.Name == "" {
		return models.ProductDTO{}, fmt.Errorf("Products.Set: %v", models.NewEmptyErr("name"))
	}

	err := c.storage.Set(ctx, product)
	if err != nil {
		return models.ProductDTO{}, fmt.Errorf("Products.Set: %v", err)
	}

	return product, nil
}

func (c Products) Get(ctx context.Context, id int) (models.ProductDTO, error) {
	//TODO: move to validateFn
	if id <= 0 {
		return models.ProductDTO{}, fmt.Errorf("Products.Get: %v", models.NewEmptyErr("id"))
	}

	product, err := c.storage.Get(ctx, id)
	if err != nil {
		return models.ProductDTO{}, fmt.Errorf("Products.Get: %v", err)
	}

	return product.ToDTO(), nil
}

func (c Products) Delete(ctx context.Context, id int) error {
	//TODO: move to validateFn
	if id <= 0 {
		return fmt.Errorf("Products.Add: %v", models.NewEmptyErr("id"))
	}

	return c.storage.Delete(ctx, id)
}
