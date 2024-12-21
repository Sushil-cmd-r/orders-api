package application

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/sushil-cmd-r/orders-api/db"
	"github.com/sushil-cmd-r/orders-api/store"
)

type App struct {
	logger *slog.Logger
	router http.Handler
	db     *db.DB
	store  *store.Store
	cfg    config
}

func New(logger *slog.Logger) *App {
	return &App{logger: logger}
}

func (a *App) Start(ctx context.Context) error {
	if err := a.loadConfig(); err != nil {
		return fmt.Errorf("failed to load config: %v", err)
	}

	if err := a.connectToDB(); err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	a.store = store.Init(a.db)
	a.loadRoutes()
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.cfg.Port),
		Handler: a.router,
	}

	done := make(chan error, 1)
	go func() {
		defer close(done)
		a.logger.Info("starting server", "port", a.cfg.Port)
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

func (a *App) connectToDB() error {
	db, err := db.Connect(a.cfg.DBUrl)
	if err != nil {
		return err
	}

	a.db = db
	return nil
}
