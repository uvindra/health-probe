package controler

import (
	"fmt"
	"health-probe/customer"
	"health-probe/models"
	"log"
	"net/http"
)

type ControlerRunner struct {
	handler          *ControlerHandler
	server           *http.Server
	name             string
	port             int
	customers        []*customer.Customer
	customerChannels []chan bool
}

type RunnerConfig struct {
	Port           int
	CustomerConfig models.CustomerConfig
	OrderSvcUrl    string
	CatalogSvcUrl  string
}

func NewRunner(cfg RunnerConfig) *ControlerRunner {
	runner := &ControlerRunner{
		handler:          NewHandler(NewControler()),
		name:             "Controler Service",
		port:             cfg.Port,
		customers:        make([]*customer.Customer, cfg.CustomerConfig.MaxCustomers),
		customerChannels: make([]chan bool, cfg.CustomerConfig.MaxCustomers),
	}

	for i := 0; i < len(runner.customers); i++ {
		runner.customers[i] = customer.NewCustomer(cfg.CatalogSvcUrl, cfg.OrderSvcUrl,
			cfg.CustomerConfig.ItemsPerOrder, cfg.CustomerConfig.QuantityPerItem)
	}

	return runner
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

	for _, c := range r.customers {
		stopChannel := make(chan bool)
		r.customerChannels = append(r.customerChannels, stopChannel)
		go c.BeginShopping(stopChannel)
	}
}

func (r *ControlerRunner) Stop() {
	if err := r.server.Close(); err != nil {
		log.Panicf("Could not stop %s: %s\n", r.name, err)
	}

	for _, c := range r.customerChannels {
		c <- true
	}
}
