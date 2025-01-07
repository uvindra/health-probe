package catalog_service

import (
	"encoding/json"
	"net/http"
)

type CatalogHandler struct {
	service *CatalogService
}

func NewInventoryRouter(service *CatalogService) *CatalogHandler {
	return &CatalogHandler{
		service: service,
	}
}

func (handler *CatalogHandler) GetItemsByCategory(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")

	if category == "" {
		http.Error(w, "Category is required", http.StatusBadRequest)
		return
	}

	items, resp := handler.service.getItemsByCategory(category)

	if resp.IsError() {
		http.Error(w, resp.Error(), resp.HttpStatus())
		return
	}

	json.NewEncoder(w).Encode(items)
}
