package application

import (
	"context"
	"log/slog"
	"net/http"
	"time"
)

type App struct {
	logger *slog.Logger
	router http.Handler
}

func New(logger *slog.Logger) *App {
	return &App{logger: logger}
}

func (a *App) Start(ctx context.Context) error {
	a.loadRoutes()

	server := &http.Server{
		Addr:    ":8080",
		Handler: a.router,
	}

	done := make(chan error, 1)
	go func() {
		defer close(done)
		a.logger.Info("starting server", "port", 8080)
		done <- server.ListenAndServe()
	}()

	select {
	case err := <-done:
		return err

	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		return server.Shutdown(timeout)
	}
}
