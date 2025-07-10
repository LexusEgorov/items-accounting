package server

import (
	"log/slog"

	"github.com/labstack/echo/v4"

	"github.com/LexusEgorov/items-accounting/internal/models"
)

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

func sendBadResponse(c echo.Context, code int, message string) error {
	return c.JSON(code, models.BadResponse{
		Message: message,
	})
}
