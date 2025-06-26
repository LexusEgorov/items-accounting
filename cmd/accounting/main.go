package main

import (
	"log/slog"
	"os"
	"os/signal"

	"github.com/LexusEgorov/items-accounting/internal/app"
	"github.com/labstack/gommon/log"
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
	log.Info("Recieved interrupt signal")
	app.Stop()
}
