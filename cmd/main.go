package main

import (
	"health-probe/customer"
	"log"
	"time"
)

func main() {
	log.Println("Starting...")
	setup()

	log.Println("Number of services to run: ", len(services))

	for _, service := range services {
		go service.Start()
	}

	go func() {
		time.Sleep(5 * time.Second)
		cfg := customer.SpawnerConfig{CustomerConfig: config.Customer, OrderSvcUrl: getOrderServiceUrl(), CatalogSvcUrl: getCatalogServiceUrl()}
		spawner := customer.NewSpawner(cfg)
		spawner.Start()
	}()

	controler.Start()
}
