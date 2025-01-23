package inventory

import (
	"fmt"
	"log"
	"net/http"
)

type InventoryServiceRunner struct {
	handler *InventoryHandler
	server  *http.Server
	name    string
	port    int
}

type RunnerConfig struct {
	Port     int
	Capacity int
	Reserve  int
}

func NewRunner(cfg RunnerConfig) *InventoryServiceRunner {
	return &InventoryServiceRunner{
		handler: NewHandler(NewService(cfg.Capacity, cfg.Reserve)),
		name:    "Inventory Service",
		port:    cfg.Port,
	}
}

func (r *InventoryServiceRunner) Start() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /items/{id}", r.handler.GetItem)
	mux.HandleFunc("GET /items", r.handler.GetItems)
	mux.HandleFunc("PATCH /items/{id}/deduct", r.handler.DeductItemQty)

	addr := fmt.Sprintf(":%d", r.port)
	r.server = &http.Server{Addr: addr, Handler: mux}

	log.Printf("Starting %s on %s\n", r.name, addr)

	if err := r.server.ListenAndServe(); err != nil {
		log.Panicf("Could not start %s: %s\n", r.name, err)
	}
}

func (r *InventoryServiceRunner) Stop() {
	if err := r.server.Close(); err != nil {
		log.Panicf("Could not stop %s: %s\n", r.name, err)
	}
}
