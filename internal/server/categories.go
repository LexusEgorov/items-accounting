package server

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/LexusEgorov/items-accounting/internal/models"
	"github.com/LexusEgorov/items-accounting/internal/utils"
	"github.com/labstack/echo/v4"
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
	errPrefix := "server.Category.Get"
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		cat.logger.Error(utils.GetError(errPrefix, err).Error())
		return utils.SendBadResponse(c, http.StatusNotFound, "id must be a number")
	}

	category, err := cat.manager.Get(c.Request().Context(), id)
	if err != nil {
		cat.logger.Error(utils.GetError(errPrefix, err).Error())
		//TODO: add check for notfoundErr
		return utils.SendBadResponse(c, http.StatusInternalServerError, "ne rabotaet")
	}

	return c.JSON(http.StatusOK, category)
}

func (cat CategoryHandler) Set(c echo.Context) error {
	errPrefix := "server.Category.Set"
	bodyReader := c.Request().Body
	defer bodyReader.Close()

	body, err := io.ReadAll(bodyReader)
	if err != nil {
		cat.logger.Error(utils.GetError(errPrefix, err).Error())
		return utils.SendBadResponse(c, http.StatusBadRequest, "error while reading body")
	}

	var category models.CategoryDTO
	err = json.Unmarshal(body, &category)
	if err != nil {
		cat.logger.Error(utils.GetError(errPrefix, err).Error())
		return utils.SendBadResponse(c, http.StatusBadRequest, "error while reading body")
	}

	updated, err := cat.manager.Set(c.Request().Context(), category)
	if err != nil {
		cat.logger.Error(utils.GetError(errPrefix, err).Error())
		//TODO: add check for notfoundErr
		return utils.SendBadResponse(c, http.StatusInternalServerError, "ne rabotaet")
	}

	return c.JSON(http.StatusOK, updated)
}

func (cat CategoryHandler) Delete(c echo.Context) error {
	errPrefix := "server.Category.Delete"
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		cat.logger.Error(utils.GetError(errPrefix, err).Error())
		return utils.SendBadResponse(c, http.StatusNotFound, "id must be a number")
	}

	err = cat.manager.Delete(c.Request().Context(), id)
	if err != nil {
		cat.logger.Error(utils.GetError(errPrefix, err).Error())
		//TODO: add check for notfoundErr
		return utils.SendBadResponse(c, http.StatusInternalServerError, "ne rabotaet")
	}

	return c.JSON(http.StatusOK, nil)
}

func (cat CategoryHandler) Add(c echo.Context) error {
	errPrefix := "server.Category.Create"
	bodyReader := c.Request().Body
	defer bodyReader.Close()

	body, err := io.ReadAll(bodyReader)
	if err != nil {
		cat.logger.Error(utils.GetError(errPrefix, err).Error())
		return utils.SendBadResponse(c, http.StatusBadRequest, "error while reading body")
	}

	var category models.CategoryDTO
	err = json.Unmarshal(body, &category)
	if err != nil {
		cat.logger.Error(utils.GetError(errPrefix, err).Error())
		return utils.SendBadResponse(c, http.StatusBadRequest, "error while reading body")
	}

	created, err := cat.manager.Add(c.Request().Context(), category.Name)
	if err != nil {
		cat.logger.Error(utils.GetError(errPrefix, err).Error())
		//TODO: add check for notfoundErr
		return utils.SendBadResponse(c, http.StatusInternalServerError, "ne rabotaet")
	}

	return c.JSON(http.StatusCreated, created)
}
