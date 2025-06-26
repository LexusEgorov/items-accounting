package utils

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/LexusEgorov/items-accounting/internal/models"
)

func GetError(prefix string, err error) error {
	if err != nil {
		return fmt.Errorf("%s: %v", prefix, err)
	}

	return nil
}

func SendBadResponse(c echo.Context, code int, message string) error {
	return c.JSON(code, models.BadResponse{
		Message: message,
	})
}
