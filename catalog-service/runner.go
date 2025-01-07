package catalog_service

import (
	"fmt"
	"log"
	"net/http"
)

type CatalogServiceRunner struct {
	router *CatalogHandler
	server *http.Server
	name   string
	port   int
}

func NewCatalogServiceRunner(port int, inventorySvcUrl string) *CatalogServiceRunner {
	return &CatalogServiceRunner{
		router: NewInventoryRouter(NewCatalogService(inventorySvcUrl)),
		name:   "Catalog Service",
		port:   port,
	}
}

func (r *CatalogServiceRunner) Start() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /items", r.router.GetItemsByCategory)

	addr := fmt.Sprintf(":%d", r.port)

	log.Printf("Starting %s on %s\n", r.name, addr)
	r.server = &http.Server{Addr: addr, Handler: mux}

	if err := r.server.ListenAndServe(); err != nil {
		log.Fatalf("Could not start %s: %s\n", r.name, err)
	}

}

func (r *CatalogServiceRunner) Stop() {
	if err := r.server.Close(); err != nil {
		log.Fatalf("Could not stop %s: %s\n", r.name, err)
	}
}
