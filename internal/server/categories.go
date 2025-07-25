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

type CategoryManager interface {
	Add(c context.Context, name string) (models.CategoryDTO, error)
	Set(c context.Context, category models.CategoryDTO) (models.CategoryDTO, error)
	Get(c context.Context, ID int) (models.CategoryDTO, error)
	Delete(c context.Context, id int) error
}

type CategoryHandler struct {
	manager CategoryManager
	logger  *slog.Logger
}

func newCategoryHandler(manager CategoryManager, logger *slog.Logger) *CategoryHandler {
	return &CategoryHandler{
		manager: manager,
		logger:  logger,
	}
}

func (cat CategoryHandler) Get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		cat.logger.Error(err.Error())
		return sendBadResponse(c, http.StatusNotFound, "id must be a number")
	}

	category, err := cat.manager.Get(c.Request().Context(), id)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return sendBadResponse(c, http.StatusNotFound, "Not found")
		}
		cat.logger.Error(err.Error())

		return sendBadResponse(c, http.StatusInternalServerError, "ne rabotaet")
	}

	return c.JSON(http.StatusOK, category)
}

func (cat CategoryHandler) Set(c echo.Context) error {
	bodyReader := c.Request().Body
	defer bodyReader.Close()

	body, err := io.ReadAll(bodyReader)
	if err != nil {
		cat.logger.Error(err.Error())
		return sendBadResponse(c, http.StatusBadRequest, "error while reading body")
	}

	var category models.CategoryDTO
	err = json.Unmarshal(body, &category)
	if err != nil {
		cat.logger.Error(err.Error())
		return sendBadResponse(c, http.StatusBadRequest, "error while reading body")
	}

	updated, err := cat.manager.Set(c.Request().Context(), category)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return sendBadResponse(c, http.StatusNotFound, "Not found")
		}

		cat.logger.Error(err.Error())
		return sendBadResponse(c, http.StatusInternalServerError, "ne rabotaet")
	}

	return c.JSON(http.StatusOK, updated)
}

func (cat CategoryHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		cat.logger.Error(err.Error())
		return sendBadResponse(c, http.StatusNotFound, "id must be a number")
	}

	err = cat.manager.Delete(c.Request().Context(), id)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return sendBadResponse(c, http.StatusNotFound, "Not found")
		}

		cat.logger.Error(err.Error())
		return sendBadResponse(c, http.StatusInternalServerError, "ne rabotaet")
	}

	return nil
}

func (cat CategoryHandler) Add(c echo.Context) error {
	bodyReader := c.Request().Body
	defer bodyReader.Close()

	body, err := io.ReadAll(bodyReader)
	if err != nil {
		cat.logger.Error(err.Error())
		return sendBadResponse(c, http.StatusBadRequest, "error while reading body")
	}

	var category models.CategoryDTO
	err = json.Unmarshal(body, &category)
	if err != nil {
		cat.logger.Error(err.Error())
		return sendBadResponse(c, http.StatusBadRequest, "error while reading body")
	}

	created, err := cat.manager.Add(c.Request().Context(), category.Name)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return sendBadResponse(c, http.StatusNotFound, "Not found")
		}

		cat.logger.Error(err.Error())
		return sendBadResponse(c, http.StatusInternalServerError, "ne rabotaet")
	}

	return c.JSON(http.StatusCreated, created)
}
