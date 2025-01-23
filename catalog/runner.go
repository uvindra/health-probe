package catalog

import (
	"fmt"
	"log"
	"net/http"
)

type CatalogServiceRunner struct {
	handler *CatalogHandler
	server  *http.Server
	name    string
	port    int
}

type RunnerConfig struct {
	Port            int
	Capacity        int
	InventorySvcUrl string
}

func NewRunner(cfg RunnerConfig) *CatalogServiceRunner {
	return &CatalogServiceRunner{
		handler: NewHandler(NewService(cfg.Capacity, cfg.InventorySvcUrl)),
		name:    "Catalog Service",
		port:    cfg.Port,
	}
}

func (r *CatalogServiceRunner) Start() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /suggestion", r.handler.GetSuggestion)

	addr := fmt.Sprintf(":%d", r.port)
	r.server = &http.Server{Addr: addr, Handler: mux}

	log.Printf("Starting %s on %s\n", r.name, addr)

	if err := r.server.ListenAndServe(); err != nil {
		log.Panicf("Could not start %s: %s\n", r.name, err)
	}
}

func (r *CatalogServiceRunner) Stop() {
	if err := r.server.Close(); err != nil {
		log.Panicf("Could not stop %s: %s\n", r.name, err)
	}
}
