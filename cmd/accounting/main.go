package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"github.com/LexusEgorov/items-accounting/internal/app"
)

func main() {
	logger := slog.Default()
	app, err := app.New(logger)

	if err != nil {
		logger.Error(err.Error())
		return
	}

	app.Run()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)

	<-stopChan
	logger.Info("Recieved interrupt signal")
	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	app.Stop(ctx)
}
