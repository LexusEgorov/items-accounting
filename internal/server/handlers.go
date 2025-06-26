package server

import "log/slog"

type handlers struct {
	categories *CategoryHandler
	products   *ProductHandler
}

func NewHandlers(category CategoryManager, product ProductManager, logger *slog.Logger) *handlers {
	return &handlers{
		categories: newCategoryHandler(category, logger),
		products:   newProductHandler(product, logger),
	}
}
