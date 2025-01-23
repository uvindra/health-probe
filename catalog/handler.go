package catalog

import (
	"encoding/json"
	"net/http"
)

type CatalogHandler struct {
	service *Catalog
}

func NewHandler(service *Catalog) *CatalogHandler {
	return &CatalogHandler{service: service}
}

func (h *CatalogHandler) GetSuggestion(w http.ResponseWriter, r *http.Request) {
	item, resp := h.service.GetSuggestion()

	if resp.IsError() {
		http.Error(w, resp.Error(), resp.HttpStatus())
		return
	}

	json.NewEncoder(w).Encode(item)
}
