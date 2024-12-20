package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"github.com/sushil-cmd-r/orders-api/application"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	app := application.New(logger)
	if err := app.Start(ctx); err != nil {
		logger.Error("failed to start application", "error", err)
	}
}
