package controler

import "net/http"

type ControlerHandler struct {
	service *Controler
}

func NewHandler(service *Controler) *ControlerHandler {
	return &ControlerHandler{
		service: service,
	}
}

func (h *ControlerHandler) GetStatus(w http.ResponseWriter, r *http.Request) {}
