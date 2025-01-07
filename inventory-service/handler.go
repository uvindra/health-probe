package inventory_service

import (
	"encoding/json"
	"net/http"
)

type InventoryHandler struct {
	service *InventoryService
}

func NewInventoryRouter(service *InventoryService) *InventoryHandler {
	return &InventoryHandler{
		service: service,
	}
}

func (router *InventoryHandler) GetItem(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	item, resp := router.service.getItem(id)

	if resp.IsError() {
		http.Error(w, resp.Error(), resp.HttpStatus())
		return
	}

	json.NewEncoder(w).Encode(item)
}

func (router *InventoryHandler) AddItem(w http.ResponseWriter, r *http.Request) {
	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp := router.service.addItem(item)
	if resp.IsError() {
		http.Error(w, resp.Error(), resp.HttpStatus())
		return
	}
	w.WriteHeader(resp.HttpStatus())
	json.NewEncoder(w).Encode(item)
}

func (router *InventoryHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp := router.service.updateItem(item)
	if resp.IsError() {
		http.Error(w, resp.Error(), resp.HttpStatus())
		return
	}

	json.NewEncoder(w).Encode(item)
}

func (router *InventoryHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	resp := router.service.deleteItem(id)

	if resp.IsError() {
		http.Error(w, resp.Error(), resp.HttpStatus())
		return
	}

	w.WriteHeader(resp.HttpStatus())
}
