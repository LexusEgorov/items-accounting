package server

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/LexusEgorov/items-accounting/internal/models"
)

type ProductManager interface {
	Add(ctx context.Context, product models.ProductDTO) (models.ProductDTO, error)
	Set(ctx context.Context, product models.ProductDTO) (models.ProductDTO, error)
	Get(ctx context.Context, id int) (models.ProductDTO, error)
	Delete(ctx context.Context, id int) error
}

type ProductHandler struct {
	manager ProductManager
	logger  *slog.Logger
}

func newProductHandler(manager ProductManager, logger *slog.Logger) *ProductHandler {
	return &ProductHandler{
		manager: manager,
		logger:  logger,
	}
}

func (p ProductHandler) Get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		p.logger.Error(err.Error())
		return sendBadResponse(c, http.StatusNotFound, "id must be a number")
	}

	product, err := p.manager.Get(c.Request().Context(), id)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return sendBadResponse(c, http.StatusNotFound, "Not found")
		}

		p.logger.Error(err.Error())
		return sendBadResponse(c, http.StatusInternalServerError, "ne rabotaet")
	}

	return c.JSON(http.StatusOK, product)
}

func (p ProductHandler) Set(c echo.Context) error {
	bodyReader := c.Request().Body
	defer bodyReader.Close()

	body, err := io.ReadAll(bodyReader)
	if err != nil {
		p.logger.Error(err.Error())
		return sendBadResponse(c, http.StatusBadRequest, "error while reading body")
	}

	var product models.ProductDTO
	err = json.Unmarshal(body, &product)
	if err != nil {
		p.logger.Error(err.Error())
		return sendBadResponse(c, http.StatusBadRequest, "error while reading body")
	}

	updated, err := p.manager.Set(c.Request().Context(), product)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return sendBadResponse(c, http.StatusNotFound, "Not found")
		}

		p.logger.Error(err.Error())
		return sendBadResponse(c, http.StatusInternalServerError, "ne rabotaet")
	}

	return c.JSON(http.StatusOK, updated)
}

func (p ProductHandler) Add(c echo.Context) error {
	bodyReader := c.Request().Body
	defer bodyReader.Close()

	body, err := io.ReadAll(bodyReader)
	if err != nil {
		p.logger.Error(err.Error())
		return sendBadResponse(c, http.StatusBadRequest, "error while reading body")
	}

	var product models.ProductDTO
	err = json.Unmarshal(body, &product)
	if err != nil {
		p.logger.Error(err.Error())
		return sendBadResponse(c, http.StatusBadRequest, "error while reading body")
	}

	created, err := p.manager.Add(c.Request().Context(), product)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return sendBadResponse(c, http.StatusNotFound, "Not found")
		}

		p.logger.Error(err.Error())
		return sendBadResponse(c, http.StatusInternalServerError, "ne rabotaet")
	}

	return c.JSON(http.StatusOK, created)
}

func (p ProductHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		p.logger.Error(err.Error())
		return sendBadResponse(c, http.StatusNotFound, "id must be a number")
	}

	err = p.manager.Delete(c.Request().Context(), id)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return sendBadResponse(c, http.StatusNotFound, "Not found")
		}

		p.logger.Error(err.Error())
		return sendBadResponse(c, http.StatusInternalServerError, "ne rabotaet")
	}

	return nil
}
