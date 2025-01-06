package main

import (
	ims "health-probe/inventory_service"
	oms "health-probe/order_service"
)

type Service interface {
	Start(port int)
	Stop()
}

func main() {
	services := make([]Service, 2)

	services = append(services, oms.NewOrderServiceRunner())
	services = append(services, ims.NewInventoryServiceRunner())

	basePort := 8080
	offset := 0
	for _, service := range services {
		service.Start(basePort + offset)
		offset++
	}

}
