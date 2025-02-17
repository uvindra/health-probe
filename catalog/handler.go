package catalog

import (
	"encoding/json"
	probe "health-probe/probe"
	"net/http"
)

type CatalogHandler struct {
	service *Catalog
}

func NewHandler(service *Catalog) *CatalogHandler {
	return &CatalogHandler{service: service}
}

func (h *CatalogHandler) GetSuggestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	item, resp := h.service.GetSuggestion()

	if resp.IsError() {
		http.Error(w, resp.Error(), resp.HttpStatus())
		return
	}

	err := json.NewEncoder(w).Encode(item)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *CatalogHandler) GetHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	probe.WriteProbes(h.service, w)
}
