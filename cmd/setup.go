package main

import (
	"fmt"
	acc "health-probe/account"
	ims "health-probe/inventory"
	"log"
	"os"

	oms "health-probe/order"

	"gopkg.in/yaml.v3"
)

const ACCOUNT_SERVICE = "AccountService"
const INVENTORY_SERVICE = "InventoryService"
const ORDER_SERVICE = "OrderService"

type Service interface {
	Start()
	Stop()
}

type ServiceConfig struct {
	name string `yaml:"name"`
	host string `yaml:"host"`
	port int    `yaml:"port"`
}

type Config struct {
	services []ServiceConfig `yaml:"services"`
}

var configLookup map[string]ServiceConfig
var serviceLookup map[string]Service

func setup() {
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Panicf("yamlFile.Get err   #%v ", err)
	}
	config := Config{}

	err = yaml.Unmarshal(yamlFile, &config)

	if err != nil {
		log.Panicf("Unmarshal: %v", err)
	}

	for _, service := range config.services {
		switch service.name {
		case ACCOUNT_SERVICE:
			configLookup[ACCOUNT_SERVICE] = service
		case INVENTORY_SERVICE:
			configLookup[INVENTORY_SERVICE] = service
		case ORDER_SERVICE:
			configLookup[ORDER_SERVICE] = service
		}
	}

	for _, service := range config.services {
		switch service.name {
		case ACCOUNT_SERVICE:
			serviceLookup[ACCOUNT_SERVICE] = acc.NewRunner(service.port)
		case INVENTORY_SERVICE:
			serviceLookup[INVENTORY_SERVICE] = ims.NewRunner(service.port)
		case ORDER_SERVICE:
			serviceLookup[ORDER_SERVICE] = oms.NewRunner(service.port, getInventoryServiceUrl())
		}
	}
}

func getInventoryServiceUrl() string {
	config, exists := configLookup[INVENTORY_SERVICE]

	if !exists {
		log.Panic("Inventory service not found in config")
	}

	return config.host + ":" + fmt.Sprint(config.port)
}
