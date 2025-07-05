package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/LexusEgorov/items-accounting/internal/models"
	"github.com/LexusEgorov/items-accounting/internal/utils"
	"github.com/labstack/echo/v4"
)

type middleware struct {
	maxResponseTime time.Duration
	logger          *slog.Logger
}

func New(logger *slog.Logger, maxResponseTime time.Duration) *middleware {
	return &middleware{
		maxResponseTime: maxResponseTime,
		logger:          logger,
	}
}

func (m middleware) WithLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		errPrefix := "middleware.WithLogging"
		timeStart := time.Now()
		err := next(c)

		if err != nil {
			m.logger.Error(utils.GetError(errPrefix, err).Error())
		}

		code := c.Response().Status
		if code >= http.StatusBadRequest && code <= http.StatusNetworkAuthenticationRequired {
			m.logger.Error("request result",
				"code", code,
				"method", c.Request().Method,
				"url", c.Request().URL,
				"duration", time.Since(timeStart))
		} else {
			m.logger.Info("request result",
				"code", code,
				"method", c.Request().Method,
				"url", c.Request().URL,
				"duration", time.Since(timeStart))
		}

		return err
	}
}

func (m middleware) WithCancel(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, cancel := context.WithTimeout(c.Request().Context(), m.maxResponseTime)
		defer cancel()

		c.SetRequest(c.Request().WithContext(ctx))

		done := make(chan error, 1)

		go func() {
			done <- next(c)
		}()

		select {
		case err := <-done:
			return err
		case <-ctx.Done():
			return c.JSON(echo.ErrInternalServerError.Code, models.BadResponse{
				Message: "request timeout",
			})
		}
	}
}

func (m middleware) WithRecover(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if r := recover(); r != nil {
				m.logger.Error("Recovered: %v", r)
				c.JSON(echo.ErrInternalServerError.Code, models.BadResponse{
					Message: "internal server error",
				})
			}
		}()

		return next(c)
	}
}
