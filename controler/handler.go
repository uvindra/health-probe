package controler

import (
	"encoding/json"
	"net/http"
)

type ControlerHandler struct {
	service *Controler
}

func NewHandler(service *Controler) *ControlerHandler {
	return &ControlerHandler{
		service: service,
	}
}

func (h *ControlerHandler) GetStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	status := h.service.GetStatus()

	err := json.NewEncoder(w).Encode(status)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
