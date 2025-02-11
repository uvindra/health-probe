package models

import "health-probe/enum"

type ServiceConfig struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type InventoryConfig struct {
	MaxItems        int `yaml:"maxItems"`
	MaxItemQuantity int `yaml:"maxItemQuantity"`
}

type CustomerConfig struct {
	MaxCustomers    int `yaml:"maxCustomers"`
	ItemsPerOrder   int `yaml:"itemsPerOrder"`
	QuantityPerItem int `yaml:"quantityPerItem"`
}

type DependencyProbe struct {
	ClientName     string `json:"clientName"`
	ServerName     string `json:"serverName"`
	TotalFailed    uint32 `json:"totalFailed"`
	TotalCompleted uint32 `json:"totalCompleted"`
}

type LocalProbe struct {
	Name           string `json:"name"`
	TotalFailed    uint32 `json:"totalFailed"`
	TotalCompleted uint32 `json:"totalCompleted"`
}

type Health struct {
	DependencyProbes []DependencyProbe `json:"dependencyProbes"`
	LocalProbes      []LocalProbe      `json:"localProbes"`
}

type Item struct {
	Id       int    `json:"Id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type OrderQty struct {
	Quantity int `json:"quantity"`
}

type Order struct {
	CustomerId string `json:"customerId"`
	Items      []Item `json:"items"`
}

type OrderTracker struct {
	Id         string
	CustomerId string
	Items      []Item
	State      enum.OrderState
}
