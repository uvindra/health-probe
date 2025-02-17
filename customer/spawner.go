package customer

import (
	"health-probe/models"
	"log"
	"sync"
)

type Spawner struct {
	customers        []*Customer
	customerChannels []chan bool
}

type SpawnerConfig struct {
	CustomerConfig models.CustomerConfig
	OrderSvcUrl    string
	CatalogSvcUrl  string
}

func NewSpawner(cfg SpawnerConfig) *Spawner {
	spawner := &Spawner{
		customers:        make([]*Customer, cfg.CustomerConfig.MaxCustomers),
		customerChannels: make([]chan bool, cfg.CustomerConfig.MaxCustomers),
	}

	for i := 0; i < len(spawner.customers); i++ {
		spawner.customers[i] = NewCustomer(cfg.CatalogSvcUrl, cfg.OrderSvcUrl,
			cfg.CustomerConfig.ItemsPerOrder, cfg.CustomerConfig.QuantityPerItem)
	}

	return spawner
}

func (s *Spawner) Start() {
	log.Println("Starting to shop")

	var wg sync.WaitGroup

	for _, c := range s.customers {
		wg.Add(1)
		stopChannel := make(chan bool)
		s.customerChannels = append(s.customerChannels, stopChannel)
		go func() {
			defer wg.Done()
			c.BeginShopping(stopChannel)
		}()
	}

	wg.Wait()

	log.Println("All customers have finished shopping")
}
