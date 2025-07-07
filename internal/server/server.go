package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/LexusEgorov/items-accounting/internal/config"
	"github.com/LexusEgorov/items-accounting/internal/middleware"
	"github.com/LexusEgorov/items-accounting/internal/utils"
	"github.com/labstack/echo/v4"
)

type Server struct {
	server *echo.Echo
	logger *slog.Logger
	config config.ServerConfig
}

func New(handlers handlers, logger *slog.Logger, config config.ServerConfig) *Server {
	server := echo.New()
	middleware := middleware.New(logger, config.MaxResponseTime)

	server.Use(middleware.WithRecover, middleware.WithLogging, middleware.WithCancel)

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
		config: config,
	}
}

func (s *Server) Run() {
	errPrefix := "server.Run"
	serverAddr := fmt.Sprintf("%s:%d", s.config.Addr, s.config.Port)
	s.logger.Info(fmt.Sprintf("server is starting on %s", serverAddr))
	if err := s.server.Start(serverAddr); err != nil {
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
