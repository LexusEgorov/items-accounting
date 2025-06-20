package server

import (
	"github.com/LexusEgorov/items-accounting/internal/models"
)

type ProductManager interface {
	Add(string) (models.Product, error)
	Get(int) (models.Product, error)
	Set(models.Product) (models.Product, error)
	Delete(int) error
}

type CategoryManager interface {
	Add(string) (models.Category, error)
	Get(int) (models.Category, error)
	Set(int, string) (models.Category, error)
	Delete(int) error
}

type handlers struct {
	categories CategoryManager
	products   ProductManager
}

// TODO: handlers for managers
func newHandlers(productManager ProductManager, categoryManager CategoryManager) *handlers {
	return &handlers{
		categories: categoryManager,
		products:   productManager,
	}
}
