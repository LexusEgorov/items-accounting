package products

import (
	"context"

	"github.com/LexusEgorov/items-accounting/internal/models"
)

type Storager interface {
	Add(ctx context.Context, product models.ProductDTO) (id int, err error)
	Get(ctx context.Context, id int) (product models.Product, err error)
	Set(ctx context.Context, product models.ProductDTO) error
	Delete(ctx context.Context, id int) error
}

type Products struct {
	storage Storager
}

func New(storage Storager) *Products {
	return &Products{
		storage: storage,
	}
}

func (c Products) Add(ctx context.Context, product models.ProductDTO) (models.ProductDTO, error) {
	//TODO: move to validateFn
	if product.CatID == 0 {
		return models.ProductDTO{}, models.NewEmptyErr("categoryId")
	}

	//TODO: move to validateFn
	if product.Name == "" {
		return models.ProductDTO{}, models.NewEmptyErr("name")
	}

	id, err := c.storage.Add(ctx, product)
	if err != nil {
		return models.ProductDTO{}, err
	}

	product.ID = id
	return product, nil
}

func (c Products) Set(ctx context.Context, product models.ProductDTO) (models.ProductDTO, error) {
	//TODO: move to validateFn
	if product.CatID <= 0 {
		return models.ProductDTO{}, models.NewEmptyErr("categoryId")
	}

	//TODO: move to validateFn
	if product.Name == "" {
		return models.ProductDTO{}, models.NewEmptyErr("name")
	}

	err := c.storage.Set(ctx, product)
	if err != nil {
		return models.ProductDTO{}, err
	}

	return product, nil
}

func (c Products) Get(ctx context.Context, id int) (models.ProductDTO, error) {
	//TODO: move to validateFn
	if id <= 0 {
		return models.ProductDTO{}, models.NewEmptyErr("id")
	}

	product, err := c.storage.Get(ctx, id)
	if err != nil {
		return models.ProductDTO{}, err
	}

	return product.ToDTO(), nil
}

func (c Products) Delete(ctx context.Context, id int) error {
	//TODO: move to validateFn
	if id <= 0 {
		return models.NewEmptyErr("id")
	}

	return c.storage.Delete(ctx, id)
}
