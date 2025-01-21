package account

import (
	"fmt"
	"log"
	"net/http"
)

type AccountRunner struct {
	handler *AccountHandler
	server  *http.Server
	name    string
	port    int
}

func NewRunner(port int) *AccountRunner {
	return &AccountRunner{
		handler: NewHandler(NewAccount()),
		name:    "Account Service",
		port:    port,
	}
}

func (r *AccountRunner) Start() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /items/{id}", r.handler.GetItem)
	mux.HandleFunc("GET /items", r.handler.GetItems)
	mux.HandleFunc("POST /items", r.handler.AddItem)
	mux.HandleFunc("DELETE /items/{id}", r.handler.RemoveItem)

	addr := fmt.Sprintf(":%d", r.port)

	log.Printf("Starting %s on %s\n", r.name, addr)
	r.server = &http.Server{Addr: addr, Handler: mux}

	if err := r.server.ListenAndServe(); err != nil {
		log.Panicf("Could not start %s: %s\n", r.name, err)
	}

}

func (r *AccountRunner) Stop() {
	if err := r.server.Close(); err != nil {
		log.Panicf("Could not stop %s: %s\n", r.name, err)
	}
}
