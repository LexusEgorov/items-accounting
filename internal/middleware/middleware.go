package middleware

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/LexusEgorov/items-accounting/internal/models"
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
		timeStart := time.Now()
		err := next(c)
		if err != nil {
			m.logger.Error(err.Error())
		}

		innerLogger := m.logger.With(
			"method", c.Request().Method,
			"url", c.Request().URL,
			"duration", time.Since(timeStart))

		code := c.Response().Status
		if code >= http.StatusBadRequest && code <= http.StatusNetworkAuthenticationRequired {
			innerLogger.Error("request result", "code", code)
		} else {
			innerLogger.Info("request result", "code", code)
		}

		return err
	}
}

func (m middleware) WithRecover(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if r := recover(); r != nil {
				m.logger.Error("recovered: %v", r)
				c.JSON(echo.ErrInternalServerError.Code, models.BadResponse{
					Message: http.StatusText(echo.ErrInternalServerError.Code),
				})
			}
		}()

		return next(c)
	}
}
