package order

import (
	"bytes"
	"encoding/json"
	"fmt"
	"health-probe/enum"
	mod "health-probe/models"
	res "health-probe/response"
	store "health-probe/store"
	"net/http"
)

type OrderStore interface {
	AddOrderTracker(order mod.Order, state enum.OrderState)
	GetOrderTracker(customerId string, orderId string) (mod.OrderTracker, bool)
	GetOrderTrackers(customerId string) []mod.OrderTracker
}

type OrderService struct {
	inventorySvcUrl string
	placedOrders    OrderStore
}

const deduct = "/items/{%d}/deduct"

func NewService(inventorySvcUrl string) *OrderService {
	return &OrderService{
		inventorySvcUrl: inventorySvcUrl,
		placedOrders:    store.NewOrderStore(),
	}
}

func (s *OrderService) placeOrder(order mod.Order) res.ServiceResponse {
	resource := s.inventorySvcUrl + deduct

	for _, item := range order.Items {
		url := fmt.Sprintf(resource, item.Id)

		orderQty := mod.OrderQty{Quantity: item.Quantity}

		payload, err := json.Marshal(orderQty)

		if err != nil {
			s.placedOrders.AddOrderTracker(order, enum.NewOrderState(enum.Failed))
			return res.NewErrorResponse(fmt.Sprintf("Error when marshaling OrderQty json: %s", err.Error()),
				http.StatusInternalServerError)
		}

		req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(payload))

		if err != nil {
			s.placedOrders.AddOrderTracker(order, enum.NewOrderState(enum.Failed))
			return res.NewErrorResponse(fmt.Sprintf("Error when building inventory request: %s", err.Error()),
				http.StatusInternalServerError)
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		defer resp.Body.Close()

		if err != nil {
			s.placedOrders.AddOrderTracker(order, enum.NewOrderState(enum.Failed))
			return res.NewErrorResponse(fmt.Sprintf("Error when calling inventory service: %s", err.Error()),
				http.StatusInternalServerError)
		}

		if resp.StatusCode != http.StatusOK {
			s.placedOrders.AddOrderTracker(order, enum.NewOrderState(enum.Failed))
			return res.NewErrorResponse(fmt.Sprintf("Unexpected status when deducting item from inventory service: %s", resp.Status),
				http.StatusInternalServerError)
		}
	}

	s.placedOrders.AddOrderTracker(order, enum.NewOrderState(enum.Successful))
	return res.NewSuccessResponse("", http.StatusCreated)
}

func (s *OrderService) fetchOrder(customerId string, orderId string) (mod.Order, res.ServiceResponse) {
	tracker, ok := s.placedOrders.GetOrderTracker(customerId, orderId)

	if !ok {
		return mod.Order{}, res.NewErrorResponse("order not found", http.StatusNotFound)
	}

	order := mod.Order{
		CustomerId: tracker.CustomerId,
		Items:      tracker.Items,
	}

	return order, res.NewSuccessResponse("", http.StatusOK)
}
