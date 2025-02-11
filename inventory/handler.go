package inventory

import (
	"encoding/json"
	mod "health-probe/models"
	"health-probe/probe"
	"net/http"
	"strconv"
)

type InventoryHandler struct {
	service *InventoryService
}

func NewHandler(service *InventoryService) *InventoryHandler {
	return &InventoryHandler{
		service: service,
	}
}

func (h *InventoryHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	items, resp := h.service.getItems()

	if resp.IsError() {
		http.Error(w, resp.Error(), resp.HttpStatus())
		return
	}

	err := json.NewEncoder(w).Encode(items)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *InventoryHandler) GetItem(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	item, resp := h.service.getItem(id)

	if resp.IsError() {
		http.Error(w, resp.Error(), resp.HttpStatus())
		return
	}

	err = json.NewEncoder(w).Encode(item)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *InventoryHandler) DeductItemQty(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var qty mod.OrderQty
	if err := json.NewDecoder(r.Body).Decode(&qty); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp := h.service.deductItemQty(id, qty.Quantity)
	if resp.IsError() {
		http.Error(w, resp.Error(), resp.HttpStatus())
		return
	}

	err = json.NewEncoder(w).Encode(qty)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *InventoryHandler) GetHealth(w http.ResponseWriter, r *http.Request) {
	probe.WriteProbes(h.service, w)
}
