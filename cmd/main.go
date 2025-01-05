package main

import (
	oms "health-probe/order_service"
)

func main() {
	order_service := oms.NewOrderServiceRunner()

	order_service.Start(8080)

}
