package app

import (
	"context"
	"log/slog"

	"github.com/LexusEgorov/items-accounting/internal/config"
	srv "github.com/LexusEgorov/items-accounting/internal/server"
	"github.com/LexusEgorov/items-accounting/internal/services/categories"
	"github.com/LexusEgorov/items-accounting/internal/services/products"
	"github.com/LexusEgorov/items-accounting/internal/storage"
	"github.com/LexusEgorov/items-accounting/internal/utils"
)

// TODO: application
type App struct {
	server *srv.Server
	logger *slog.Logger
}

func New(logger *slog.Logger, config *config.Config) (*App, error) {
	errPrefix := "app.New"
	db, err := storage.NewDB(config.DB)

	if err != nil {
		return nil, utils.GetError(errPrefix, err)
	}

	categoryStorage, err := storage.NewCategories(db)

	if err != nil {
		return nil, utils.GetError(errPrefix, err)
	}

	productStorage, err := storage.NewProducts(db)

	if err != nil {
		return nil, utils.GetError(errPrefix, err)
	}

	categoryManager := categories.New(categoryStorage)
	productManager := products.New(productStorage)

	handlers := srv.NewHandlers(categoryManager, productManager, logger)
	server := srv.New(*handlers, logger, config.Server)

	return &App{
		server: server,
		logger: logger,
	}, nil
}

func (a App) Run() {
	a.logger.Info("Starting app...")
	go a.server.Run()
}

func (a App) Stop(ctx context.Context) {
	errPrefix := "app.Stop"
	a.logger.Info("Stopping app...")

	doneCh := make(chan error)
	go func() {
		doneCh <- a.server.Stop(ctx)
	}()

	select {
	case err := <-doneCh:
		if err != nil {
			a.logger.Error(utils.GetError(errPrefix, err).Error())
			return
		}

		a.logger.Info("App has been stopped gracefully")

	case <-ctx.Done():
		a.logger.Warn("App stopped forced")
	}
}
