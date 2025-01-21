package models

import "health-probe/enum"

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
