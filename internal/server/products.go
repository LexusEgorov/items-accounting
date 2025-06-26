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

func (p ProductHandler) Get(ctx echo.Context) error {
	errPrefix := "handler.Product.Get"
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		p.logger.Error(utils.GetError(errPrefix, err).Error())
		return utils.SendBadResponse(ctx, http.StatusNotFound, "id must be a number")
	}

	product, err := p.manager.Get(context.TODO(), id)

	if err != nil {
		p.logger.Error(utils.GetError(errPrefix, err).Error())
		//TODO: add check for notfoundErr
		return utils.SendBadResponse(ctx, http.StatusInternalServerError, "ne rabotaet")
	}

	return ctx.JSON(http.StatusOK, product)
}

func (p ProductHandler) Set(ctx echo.Context) error {
	errPrefix := "handler.Product.Set"
	bodyReader := ctx.Request().Body
	defer bodyReader.Close()

	body, err := io.ReadAll(bodyReader)

	if err != nil {
		p.logger.Error(utils.GetError(errPrefix, err).Error())
		return utils.SendBadResponse(ctx, http.StatusBadRequest, "error while reading body")
	}

	var product models.ProductDTO
	err = json.Unmarshal(body, &product)

	if err != nil {
		p.logger.Error(utils.GetError(errPrefix, err).Error())
		return utils.SendBadResponse(ctx, http.StatusBadRequest, "error while reading body")
	}

	updated, err := p.manager.Set(context.TODO(), product)

	if err != nil {
		p.logger.Error(utils.GetError(errPrefix, err).Error())
		//TODO: add check for notfoundErr
		return utils.SendBadResponse(ctx, http.StatusInternalServerError, "ne rabotaet")
	}

	return ctx.JSON(http.StatusOK, updated)
}

func (p ProductHandler) Add(ctx echo.Context) error {
	errPrefix := "handler.Product.Add"
	bodyReader := ctx.Request().Body
	defer bodyReader.Close()

	body, err := io.ReadAll(bodyReader)

	if err != nil {
		p.logger.Error(utils.GetError(errPrefix, err).Error())
		return utils.SendBadResponse(ctx, http.StatusBadRequest, "error while reading body")
	}

	var product models.ProductDTO
	err = json.Unmarshal(body, &product)

	if err != nil {
		p.logger.Error(utils.GetError(errPrefix, err).Error())
		return utils.SendBadResponse(ctx, http.StatusBadRequest, "error while reading body")
	}

	created, err := p.manager.Add(context.TODO(), product)

	if err != nil {
		p.logger.Error(utils.GetError(errPrefix, err).Error())
		//TODO: add check for notfoundErr
		return utils.SendBadResponse(ctx, http.StatusInternalServerError, "ne rabotaet")
	}

	return ctx.JSON(http.StatusOK, created)
}

func (p ProductHandler) Delete(ctx echo.Context) error {
	errPrefix := "handler.Product.Delete"
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		p.logger.Error(utils.GetError(errPrefix, err).Error())
		return utils.SendBadResponse(ctx, http.StatusNotFound, "id must be a number")
	}

	err = p.manager.Delete(context.TODO(), id)

	if err != nil {
		p.logger.Error(utils.GetError(errPrefix, err).Error())
		//TODO: add check for notfoundErr
		return utils.SendBadResponse(ctx, http.StatusInternalServerError, "ne rabotaet")
	}

	return ctx.JSON(http.StatusOK, nil)
}
