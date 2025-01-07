package inventory_service

import (
	"fmt"
	"log"
	"net/http"
)

type InventoryServiceRunner struct {
	router *InventoryHandler
	server *http.Server
	name   string
	port   int
}

func NewInventoryServiceRunner(port int) *InventoryServiceRunner {
	return &InventoryServiceRunner{
		router: NewInventoryRouter(NewInventoryService()),
		name:   "Inventory Service",
		port:   port,
	}
}

func (r *InventoryServiceRunner) Start() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /items", r.router.AddItem)
	mux.HandleFunc("GET /items/{id}", r.router.GetItem)
	mux.HandleFunc("PUT /items", r.router.UpdateItem)
	mux.HandleFunc("DELETE /items/{id}", r.router.DeleteItem)

	addr := fmt.Sprintf(":%d", r.port)

	log.Printf("Starting %s on %s\n", r.name, addr)
	r.server = &http.Server{Addr: addr, Handler: mux}

	if err := r.server.ListenAndServe(); err != nil {
		log.Fatalf("Could not start %s: %s\n", r.name, err)
	}

}

func (r *InventoryServiceRunner) Stop() {
	if err := r.server.Close(); err != nil {
		log.Fatalf("Could not stop %s: %s\n", r.name, err)
	}
}
