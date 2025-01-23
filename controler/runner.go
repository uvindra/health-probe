package controler

import (
	"fmt"
	"log"
	"net/http"
)

type ControlerRunner struct {
	handler *ControlerHandler
	server  *http.Server
	name    string
	port    int
}

func NewRunner(port int) *ControlerRunner {
	return &ControlerRunner{
		handler: NewHandler(NewControler()),
		name:    "Controler Service",
		port:    port,
	}
}

func (r *ControlerRunner) Start() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /status", r.handler.GetStatus)

	addr := fmt.Sprintf(":%d", r.port)

	r.server = &http.Server{Addr: addr, Handler: mux}

	log.Printf("Starting %s on %s\n", r.name, addr)

	if err := r.server.ListenAndServe(); err != nil {
		log.Panicf("Could not start %s: %s\n", r.name, err)
	}

}

func (r *ControlerRunner) Stop() {
	if err := r.server.Close(); err != nil {
		log.Panicf("Could not stop %s: %s\n", r.name, err)
	}
}
