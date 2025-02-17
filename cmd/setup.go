package main

import (
	"embed"
	"fmt"
	cat "health-probe/catalog"
	ctl "health-probe/controler"
	ims "health-probe/inventory"
	"health-probe/models"
	"log"

	oms "health-probe/order"

	"gopkg.in/yaml.v3"
)

const CONTROLER_SERVICE = "ControlerService"
const CATALOG_SERVICE = "CatalogService"
const INVENTORY_SERVICE = "InventoryService"
const ORDER_SERVICE = "OrderService"

type Service interface {
	Start()
	Stop()
}

type Config struct {
	Services  []models.ServiceConfig `yaml:"services"`
	Inventory models.InventoryConfig `yaml:"inventory"`
	Customer  models.CustomerConfig  `yaml:"customer"`
}

var serviceConfigLookup = make(map[string]models.ServiceConfig)
var services = make(map[string]Service)
var controler Service
var config Config

//go:embed config.yaml
var content embed.FS

func setup() {

	config = *readYaml()

	for _, service := range config.Services {
		switch service.Name {
		case CATALOG_SERVICE:
			serviceConfigLookup[CATALOG_SERVICE] = service
		case INVENTORY_SERVICE:
			serviceConfigLookup[INVENTORY_SERVICE] = service
		case ORDER_SERVICE:
			serviceConfigLookup[ORDER_SERVICE] = service
		}
	}

	for _, service := range config.Services {
		switch service.Name {
		case CONTROLER_SERVICE:
			cfg := ctl.RunnerConfig{Port: service.Port, CustomerConfig: config.Customer, OrderSvcUrl: getOrderServiceUrl(),
				CatalogSvcUrl: getCatalogServiceUrl(), InventorySvcUrl: getInventoryServiceUrl()}
			controler = ctl.NewRunner(cfg)
		case CATALOG_SERVICE:
			cfg := cat.RunnerConfig{Port: service.Port, Capacity: config.Inventory.MaxItems, InventorySvcUrl: getInventoryServiceUrl()}
			services[CATALOG_SERVICE] = cat.NewRunner(cfg)
		case INVENTORY_SERVICE:
			cfg := ims.RunnerConfig{Port: service.Port, Capacity: config.Inventory.MaxItems, Reserve: config.Inventory.MaxItemQuantity}
			services[INVENTORY_SERVICE] = ims.NewRunner(cfg)
		case ORDER_SERVICE:
			services[ORDER_SERVICE] = oms.NewRunner(service.Port, getInventoryServiceUrl())
		}
	}
}

func readYaml() *Config {
	yamlFile, err := content.ReadFile("config.yaml")

	if err != nil {
		log.Panicf("yamlFile.Get err   #%v ", err)
	}
	var config Config

	err = yaml.Unmarshal(yamlFile, &config)

	if err != nil {
		log.Panicf("Unmarshal: %v", err)
	}

	return &config
}

func getInventoryServiceUrl() string {
	config, exists := serviceConfigLookup[INVENTORY_SERVICE]

	if !exists {
		log.Panic("Inventory service not found in config")
	}

	return config.Host + ":" + fmt.Sprint(config.Port)
}

func getCatalogServiceUrl() string {
	config, exists := serviceConfigLookup[CATALOG_SERVICE]

	if !exists {
		log.Panic("Catalog service not found in config")
	}

	return config.Host + ":" + fmt.Sprint(config.Port)
}

func getOrderServiceUrl() string {
	config, exists := serviceConfigLookup[ORDER_SERVICE]

	if !exists {
		log.Panic("Order service not found in config")
	}

	return config.Host + ":" + fmt.Sprint(config.Port)
}
