package utils

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/LexusEgorov/items-accounting/internal/models"
)

func SendBadResponse(c echo.Context, code int, message string) error {
	return c.JSON(code, models.BadResponse{
		Message: message,
	})
}

func GetConnStr(user, password, name string) string {
	return fmt.Sprintf("postgres://%s:%s@db:5432/%s?sslmode=disable", user, password, name)
}
