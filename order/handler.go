package order

import (
	"log/slog"
	"net/http"

	"github.com/sushil-cmd-r/orders-api/store"
)

type Handler struct {
	Logger *slog.Logger
	Store  *store.Store
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("create order")
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("list orders")
}

func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("get order by id")
}

func (h *Handler) UpdateById(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("update order by id")
}

func (h *Handler) DeleteById(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("delete order by id")
}
