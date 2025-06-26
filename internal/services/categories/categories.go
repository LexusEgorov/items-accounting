package categories

import (
	"context"

	"github.com/LexusEgorov/items-accounting/internal/models"
	"github.com/LexusEgorov/items-accounting/internal/utils"
)

type Storager interface {
	Add(ctx context.Context, name string) (id int, err error)
	Get(ctx context.Context, id int) (category models.Category, err error)
	Set(ctx context.Context, id int, name string) error
	Delete(ctx context.Context, id int) error
}

type Categories struct {
	storage Storager
}

func New(storage Storager) *Categories {
	return &Categories{
		storage: storage,
	}
}

func (c Categories) Add(ctx context.Context, name string) (models.CategoryDTO, error) {
	errPrefix := "service.Categories.Add"

	//TODO: move to validateFn
	if name == "" {
		return models.CategoryDTO{}, utils.GetError(errPrefix, models.NewEmptyErr("name"))
	}

	id, err := c.storage.Add(ctx, name)

	if err != nil {
		return models.CategoryDTO{}, utils.GetError(errPrefix, err)
	}

	return models.CategoryDTO{
		ID:   id,
		Name: name,
	}, nil
}

func (c Categories) Set(ctx context.Context, category models.CategoryDTO) (models.CategoryDTO, error) {
	errPrefix := "service.Categories.Set"

	//TODO: move to validateFn
	if category.Name == "" {
		return models.CategoryDTO{}, utils.GetError(errPrefix, models.NewEmptyErr("name"))
	}

	if category.ID == 0 {
		return models.CategoryDTO{}, utils.GetError(errPrefix, models.NewEmptyErr("id"))
	}

	err := c.storage.Set(ctx, category.ID, category.Name)

	if err != nil {
		return models.CategoryDTO{}, utils.GetError(errPrefix, err)
	}

	return category, nil
}

func (c Categories) Get(ctx context.Context, ID int) (models.CategoryDTO, error) {
	errPrefix := "service.Categories.Get"
	category, err := c.storage.Get(ctx, ID)

	if err != nil {
		return models.CategoryDTO{}, utils.GetError(errPrefix, err)
	}

	return category.ToDTO(), nil
}

func (c Categories) Delete(ctx context.Context, id int) error {
	errPrefix := "service.Categories.Delete"

	//TODO: move to validateFn
	if id <= 0 {
		return utils.GetError(errPrefix, models.NewEmptyErr("id"))
	}

	return utils.GetError(errPrefix, c.storage.Delete(ctx, id))
}
