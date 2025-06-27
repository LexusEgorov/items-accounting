package server

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/LexusEgorov/items-accounting/internal/utils"
	"github.com/labstack/echo/v4"
)

type Server struct {
	server *echo.Echo
	logger *slog.Logger
}

func New(handlers handlers, logger *slog.Logger) *Server {
	server := echo.New()

	categoryGroup := server.Group("categories")
	productGroup := server.Group("products")

	categoryGroup.GET("/:id", handlers.categories.Get)
	categoryGroup.DELETE("/:id", handlers.categories.Delete)
	categoryGroup.POST("/create", handlers.categories.Add)
	categoryGroup.POST("/update", handlers.categories.Set)

	productGroup.GET("/:id", handlers.products.Get)
	productGroup.DELETE("/:id", handlers.products.Delete)
	productGroup.POST("/create", handlers.products.Add)
	productGroup.POST("/update", handlers.products.Set)

	return &Server{
		server: server,
		logger: logger,
	}
}

func (s *Server) Run() {
	errPrefix := "server.Run"
	s.logger.Info("server starting on localhost:8080")
	if err := s.server.Start("0.0.0.0:8080"); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			s.logger.Error(utils.GetError(errPrefix, err).Error())
		}
	}
}

func (s *Server) Stop(ctx context.Context) error {
	errPrefix := "server.Stop"
	s.logger.Info("stopping server...")
	err := s.server.Shutdown(ctx)

	if err != nil {
		s.logger.Error(utils.GetError(errPrefix, err).Error())
	}

	return err
}
