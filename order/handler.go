package order

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/sushil-cmd-r/orders-api/store"
	"github.com/sushil-cmd-r/orders-api/store/model"
)

type Handler struct {
	Logger *slog.Logger
	Store  *store.Store
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("create order")
	var body struct {
		CustomerId uuid.UUID        `json:"customer_id"`
		LineItems  []model.LineItem `json:"line_items"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	order := &model.Order{
		CustomerId: body.CustomerId,
		LineItems:  body.LineItems,
	}

	err := h.Store.Orders.Insert(r.Context(), order)
	if err != nil {
		h.Logger.Error("create order", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("list orders")
}

func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("get order by id")
	id := r.PathValue("id")

	orderId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		h.Logger.Error("invalid order id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	order, err := h.Store.Orders.SelectById(r.Context(), orderId)
	if errors.Is(store.ErrNotExist, err) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		h.Logger.Error("select order", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(order)
	if err != nil {
		h.Logger.Error("select order", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func (h *Handler) UpdateById(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("update order by id")
}

func (h *Handler) DeleteById(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("delete order by id")
	id := r.PathValue("id")

	orderId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		h.Logger.Error("invalid order id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.Store.Orders.DeleteById(r.Context(), orderId); err != nil {
		h.Logger.Error("delete order", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
