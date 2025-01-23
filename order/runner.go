package order

import (
	"fmt"
	"log"
	"net/http"
)

type OrderServiceRunner struct {
	handler *OrderHandler
	server  *http.Server
	name    string
	port    int
}

func NewRunner(port int, inventorySvcUrl string) *OrderServiceRunner {
	return &OrderServiceRunner{
		handler: NewHandler(NewService(inventorySvcUrl)),
		name:    "Order Service",
		port:    port,
	}
}

func (r *OrderServiceRunner) Start() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /orders", r.handler.CreateOrder)
	mux.HandleFunc("GET /orders/{id}", r.handler.GetOrder)

	addr := fmt.Sprintf(":%d", r.port)
	r.server = &http.Server{Addr: addr, Handler: mux}

	log.Printf("Starting %s on %s\n", r.name, addr)

	if err := r.server.ListenAndServe(); err != nil {
		log.Panicf("Could not start %s: %s\n", r.name, err)
	}
}

func (r *OrderServiceRunner) Stop() {
	if err := r.server.Close(); err != nil {
		log.Panicf("Could not stop %s: %s\n", r.name, err)
	}
}
