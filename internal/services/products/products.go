package products

import (
	"context"

	"github.com/LexusEgorov/items-accounting/internal/models"
	"github.com/LexusEgorov/items-accounting/internal/utils"
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
	errPrefix := "service.Products.Add"

	//TODO: move to validateFn
	if product.CatID == 0 {
		return models.ProductDTO{}, utils.GetError(errPrefix, models.NewEmptyErr("categoryId"))
	}

	//TODO: move to validateFn
	if product.Name == "" {
		return models.ProductDTO{}, utils.GetError(errPrefix, models.NewEmptyErr("name"))
	}

	id, err := c.storage.Add(ctx, product)
	if err != nil {
		return models.ProductDTO{}, utils.GetError(errPrefix, err)
	}

	product.ID = id
	return product, nil
}

func (c Products) Set(ctx context.Context, product models.ProductDTO) (models.ProductDTO, error) {
	errPrefix := "service.Products.Set"

	//TODO: move to validateFn
	if product.CatID <= 0 {
		return models.ProductDTO{}, utils.GetError(errPrefix, models.NewEmptyErr("categoryId"))
	}

	//TODO: move to validateFn
	if product.Name == "" {
		return models.ProductDTO{}, utils.GetError(errPrefix, models.NewEmptyErr("name"))
	}

	err := c.storage.Set(ctx, product)
	if err != nil {
		return models.ProductDTO{}, utils.GetError(errPrefix, err)
	}

	return product, nil
}

func (c Products) Get(ctx context.Context, id int) (models.ProductDTO, error) {
	errPrefix := "service.Products.Get"

	//TODO: move to validateFn
	if id <= 0 {
		return models.ProductDTO{}, utils.GetError(errPrefix, models.NewEmptyErr("id"))
	}

	product, err := c.storage.Get(ctx, id)
	if err != nil {
		return models.ProductDTO{}, utils.GetError(errPrefix, err)
	}

	return product.ToDTO(), nil
}

func (c Products) Delete(ctx context.Context, id int) error {
	errPrefix := "service.Products.Delete"

	//TODO: move to validateFn
	if id <= 0 {
		return utils.GetError(errPrefix, models.NewEmptyErr("id"))
	}

	return utils.GetError(errPrefix, c.storage.Delete(ctx, id))
}
