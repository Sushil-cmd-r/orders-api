package application

import (
	"net/http"

	"github.com/sushil-cmd-r/orders-api/order"
)

func (a *App) loadRoutes() {
	router := http.NewServeMux()

	router.HandleFunc("GET /health", check)

	// order routes
	orderHandler := order.Handler{Logger: a.logger, Store: a.store}
	router.HandleFunc("POST /orders", orderHandler.Create)
	router.HandleFunc("GET /orders", orderHandler.List)
	router.HandleFunc("GET /orders/{id}", orderHandler.GetById)
	router.HandleFunc("PUT /orders/{id}", orderHandler.UpdateById)
	router.HandleFunc("DELETE /orders/{id}", orderHandler.DeleteById)

	a.router = router
}

func check(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
