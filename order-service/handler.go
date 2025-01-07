package order_service

import (
	"encoding/json"
	"net/http"
)

type OrderHandler struct {
	service *OrderService
}

func NewOrderRouter(service *OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp := h.service.placeOrder(order)
	if resp.IsError() {
		http.Error(w, resp.Error(), resp.HttpStatus())
		return
	}
	w.WriteHeader(resp.HttpStatus())
	json.NewEncoder(w).Encode(order)
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	order, resp := h.service.fetchOrder(id)

	if resp.IsError() {
		http.Error(w, resp.Error(), resp.HttpStatus())
		return
	}

	json.NewEncoder(w).Encode(order)
}
