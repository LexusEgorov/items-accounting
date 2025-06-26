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
	Add(ctx context.Context, name string) (models.CategoryDTO, error)
	Set(ctx context.Context, category models.CategoryDTO) (models.CategoryDTO, error)
	Get(ctx context.Context, ID int) (models.CategoryDTO, error)
	Delete(ctx context.Context, id int) error
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

func (c CategoryHandler) Get(ctx echo.Context) error {
	errPrefix := "handler.Category.Get"
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		c.logger.Error(utils.GetError(errPrefix, err).Error())
		return utils.SendBadResponse(ctx, http.StatusNotFound, "id must be a number")
	}

	category, err := c.manager.Get(context.TODO(), id)

	if err != nil {
		c.logger.Error(utils.GetError(errPrefix, err).Error())
		//TODO: add check for notfoundErr
		return utils.SendBadResponse(ctx, http.StatusInternalServerError, "ne rabotaet")
	}

	return ctx.JSON(http.StatusOK, category)
}

func (c CategoryHandler) Set(ctx echo.Context) error {
	errPrefix := "handler.Category.Set"
	bodyReader := ctx.Request().Body
	defer bodyReader.Close()

	body, err := io.ReadAll(bodyReader)

	if err != nil {
		c.logger.Error(utils.GetError(errPrefix, err).Error())
		return utils.SendBadResponse(ctx, http.StatusBadRequest, "error while reading body")
	}

	var category models.CategoryDTO
	err = json.Unmarshal(body, &category)

	if err != nil {
		c.logger.Error(utils.GetError(errPrefix, err).Error())
		return utils.SendBadResponse(ctx, http.StatusBadRequest, "error while reading body")
	}

	updated, err := c.manager.Set(context.TODO(), category)

	if err != nil {
		c.logger.Error(utils.GetError(errPrefix, err).Error())
		//TODO: add check for notfoundErr
		return utils.SendBadResponse(ctx, http.StatusInternalServerError, "ne rabotaet")
	}

	return ctx.JSON(http.StatusOK, updated)
}

func (c CategoryHandler) Delete(ctx echo.Context) error {
	errPrefix := "handler.Category.Delete"
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		c.logger.Error(utils.GetError(errPrefix, err).Error())
		return utils.SendBadResponse(ctx, http.StatusNotFound, "id must be a number")
	}

	err = c.manager.Delete(context.TODO(), id)

	if err != nil {
		c.logger.Error(utils.GetError(errPrefix, err).Error())
		//TODO: add check for notfoundErr
		return utils.SendBadResponse(ctx, http.StatusInternalServerError, "ne rabotaet")
	}

	return ctx.JSON(http.StatusOK, nil)
}

func (c CategoryHandler) Add(ctx echo.Context) error {
	errPrefix := "handler.Category.Create"
	bodyReader := ctx.Request().Body
	defer bodyReader.Close()

	body, err := io.ReadAll(bodyReader)

	if err != nil {
		c.logger.Error(utils.GetError(errPrefix, err).Error())
		return utils.SendBadResponse(ctx, http.StatusBadRequest, "error while reading body")
	}

	var category models.CategoryDTO
	err = json.Unmarshal(body, &category)

	if err != nil {
		c.logger.Error(utils.GetError(errPrefix, err).Error())
		return utils.SendBadResponse(ctx, http.StatusBadRequest, "error while reading body")
	}

	created, err := c.manager.Add(context.TODO(), category.Name)

	if err != nil {
		c.logger.Error(utils.GetError(errPrefix, err).Error())
		//TODO: add check for notfoundErr
		return utils.SendBadResponse(ctx, http.StatusInternalServerError, "ne rabotaet")
	}

	return ctx.JSON(http.StatusCreated, created)
}
