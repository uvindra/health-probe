package order_service

import (
	"fmt"
	"log"
	"net/http"
)

type OrderServiceRunner struct {
	router *OrderHandler
	server *http.Server
	name   string
}

func NewOrderServiceRunner() *OrderServiceRunner {
	return &OrderServiceRunner{
		router: NewOrderRouter(NewOrderService()),
		name:   "Order Service",
	}
}

func (r *OrderServiceRunner) Start(port int) {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /orders", r.router.CreateOrder)
	mux.HandleFunc("GET /orders/{id}", r.router.GetOrder)

	addr := fmt.Sprintf(":%d", port)

	log.Printf("Starting %s on %s\n", r.name, addr)
	r.server = &http.Server{Addr: addr, Handler: mux}

	if err := r.server.ListenAndServe(); err != nil {
		log.Fatalf("Could not start %s: %s\n", r.name, err)
	}

}

func (r *OrderServiceRunner) Stop() {
	if err := r.server.Close(); err != nil {
		log.Fatalf("Could not stop %s: %s\n", r.name, err)
	}
}
