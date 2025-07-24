package categories

import (
	"context"
	"fmt"

	"github.com/LexusEgorov/items-accounting/internal/models"
)

type CategoryRepository interface {
	Add(ctx context.Context, name string) (id int, err error)
	Get(ctx context.Context, id int) (category models.Category, err error)
	Set(ctx context.Context, id int, name string) error
	Delete(ctx context.Context, id int) error
}

type Categories struct {
	storage CategoryRepository
}

func New(storage CategoryRepository) *Categories {
	return &Categories{
		storage: storage,
	}
}

func (c Categories) Add(ctx context.Context, name string) (models.CategoryDTO, error) {
	//TODO: move to validateFn
	if name == "" {
		return models.CategoryDTO{}, fmt.Errorf("Categories.Add: %v", models.NewEmptyErr("name"))
	}

	id, err := c.storage.Add(ctx, name)
	if err != nil {
		return models.CategoryDTO{}, fmt.Errorf("Categories.Add: %v", err)
	}

	return models.CategoryDTO{
		ID:   id,
		Name: name,
	}, nil
}

func (c Categories) Set(ctx context.Context, category models.CategoryDTO) (models.CategoryDTO, error) {
	//TODO: move to validateFn
	if category.Name == "" {
		return models.CategoryDTO{}, fmt.Errorf("Categories.Set: %v", models.NewEmptyErr("name"))
	}

	if category.ID == 0 {
		return models.CategoryDTO{}, fmt.Errorf("Categories.Set: %v", models.NewEmptyErr("id"))
	}

	err := c.storage.Set(ctx, category.ID, category.Name)
	if err != nil {
		return models.CategoryDTO{}, fmt.Errorf("Categories.Set: %v", err)
	}

	return category, nil
}

func (c Categories) Get(ctx context.Context, ID int) (models.CategoryDTO, error) {
	category, err := c.storage.Get(ctx, ID)
	if err != nil {
		return models.CategoryDTO{}, fmt.Errorf("Categories.Get: %v", err)
	}

	return category.ToDTO(), nil
}

func (c Categories) Delete(ctx context.Context, id int) error {

	//TODO: move to validateFn
	if id <= 0 {
		return fmt.Errorf("Categories.Delete: %v", models.NewEmptyErr("id"))
	}

	return c.storage.Delete(ctx, id)
}
