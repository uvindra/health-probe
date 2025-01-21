package order

import (
	"encoding/json"
	mod "health-probe/models"
	"net/http"
)

type OrderHandler struct {
	service *OrderService
}

func NewHandler(service *OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order mod.Order
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
	customerId := r.Header.Get("customerId")

	order, resp := h.service.fetchOrder(customerId, id)

	if resp.IsError() {
		http.Error(w, resp.Error(), resp.HttpStatus())
		return
	}

	json.NewEncoder(w).Encode(order)
}
