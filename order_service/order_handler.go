package order_service

import (
	"encoding/json"
	"net/http"
)

type OrderHandler struct {
	service *OrderService
}

func NewOrderHandler(service *OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := h.service.PlaceOrder(order)
	if err != nil {
		http.Error(w, err.Error(), err.HttpStatus())
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	order, err := h.service.FetchOrder(id)

	if err != nil {
		http.Error(w, err.Error(), err.HttpStatus())
		return
	}

	json.NewEncoder(w).Encode(order)
}
