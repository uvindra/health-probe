package account

import (
	"encoding/json"
	mod "health-probe/models"
	"net/http"
	"strconv"
)

type AccountHandler struct {
	service *Account
}

func NewHandler(service *Account) *AccountHandler {
	return &AccountHandler{
		service: service,
	}
}

func (h *AccountHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	items := h.service.GetItems()
	json.NewEncoder(w).Encode(items)
}

func (h *AccountHandler) GetItem(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	item, ok := h.service.GetItem(id)

	if !ok {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(item)
}

func (h *AccountHandler) AddItem(w http.ResponseWriter, r *http.Request) {
	var item mod.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.service.AddItem(item)
	w.WriteHeader(http.StatusCreated)
}

func (h *AccountHandler) RemoveItem(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	ok := h.service.RemoveItem(id)

	if !ok {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
