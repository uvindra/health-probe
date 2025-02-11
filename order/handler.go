package order

import (
	"encoding/json"
	mod "health-probe/models"
	"health-probe/probe"
	"net/http"
)

type OrderHandler struct {
	service *OrderService
	probe   probe.DependencyProbe
}

func NewHandler(service *OrderService) *OrderHandler {
	return &OrderHandler{
		service: service,
		probe:   probe.DependencyProbe{},
	}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order mod.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		h.probe.IncrementErrorCount()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := h.service.placeOrder(order)
	if resp.IsError() {
		h.probe.IncrementErrorCount()
		http.Error(w, resp.Error(), resp.HttpStatus())
		return
	}
	w.WriteHeader(resp.HttpStatus())
	err := json.NewEncoder(w).Encode(order)

	if err != nil {
		h.probe.IncrementErrorCount()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.probe.IncrementSuccessCount()
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	customerId := r.Header.Get("customerId")

	order, resp := h.service.fetchOrder(customerId, id)

	if resp.IsError() {
		h.probe.IncrementErrorCount()
		http.Error(w, resp.Error(), resp.HttpStatus())
		return
	}

	err := json.NewEncoder(w).Encode(order)

	if err != nil {
		h.probe.IncrementErrorCount()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.probe.IncrementSuccessCount()
}

func (h *OrderHandler) GetHealth(w http.ResponseWriter, r *http.Request) {
	probe.WriteProbes(h.service, w)
}
