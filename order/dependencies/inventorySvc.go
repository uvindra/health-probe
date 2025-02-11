package dependencies

import (
	"bytes"
	"encoding/json"
	"fmt"
	mod "health-probe/models"
	"health-probe/probe"
	res "health-probe/response"
	"log"
	"net/http"
)

type InventorySvc struct {
	inventorySvcUrl string
	probe           *probe.DependencyProbe
}

const deduct = "/items/{%d}/deduct"

func NewInventorySvc(inventorySvcUrl string) *InventorySvc {
	probe := probe.NewDependencyProbe("OrderSvc", "InventorySvc")
	return &InventorySvc{inventorySvcUrl: inventorySvcUrl, probe: probe}
}

func (i *InventorySvc) DeductQuantity(order mod.Order) res.ServiceResponse {
	resource := i.inventorySvcUrl + deduct

	for _, item := range order.Items {
		url := fmt.Sprintf(resource, item.Id)

		orderQty := mod.OrderQty{Quantity: item.Quantity}

		payload, err := json.Marshal(orderQty)

		if err != nil {
			log.Panicf("Error when marshaling OrderQty json: %s", err.Error())
		}

		req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(payload))

		if err != nil {
			log.Panicf("Error when building inventory request: %s", err.Error())
		}

		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			return res.NewErrorResponse(fmt.Sprintf("Error when calling inventory service: %s", err.Error()),
				http.StatusInternalServerError, i.probe.BaseProbe)
		}

		if resp.StatusCode != http.StatusOK {
			return res.NewErrorResponse(fmt.Sprintf("Unexpected status when deducting item from inventory service: %s", resp.Status),
				http.StatusInternalServerError, i.probe.BaseProbe)
		}
	}

	return res.NewSuccessResponse("", http.StatusCreated, i.probe.BaseProbe)
}

func (i *InventorySvc) GetProbe() *probe.DependencyProbe {
	return i.probe
}
