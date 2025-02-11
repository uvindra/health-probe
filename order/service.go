package order

import (
	"health-probe/enum"
	mod "health-probe/models"
	intern "health-probe/order/dependencies"
	"health-probe/probe"
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
	inventorySvc *intern.InventorySvc
	placedOrders OrderStore
	probe        *probe.LocalProbe
}

func NewService(inventorySvcUrl string) *OrderService {
	inventorySvc := intern.NewInventorySvc(inventorySvcUrl)
	return &OrderService{
		inventorySvc: inventorySvc,
		placedOrders: store.NewOrderStore(),
		probe:        probe.NewLocalProbe("OrderSvc"),
	}
}

func (s *OrderService) placeOrder(order mod.Order) res.ServiceResponse {
	res := s.inventorySvc.DeductQuantity(order)

	if res.IsError() {
		s.placedOrders.AddOrderTracker(order, enum.NewOrderState(enum.Failed))
	} else {
		s.placedOrders.AddOrderTracker(order, enum.NewOrderState(enum.Successful))
	}

	return res
}

func (s *OrderService) fetchOrder(customerId string, orderId string) (mod.Order, res.ServiceResponse) {
	tracker, ok := s.placedOrders.GetOrderTracker(customerId, orderId)

	if !ok {
		return mod.Order{}, res.NewErrorResponse("order not found", http.StatusNotFound, s.probe.BaseProbe)
	}

	order := mod.Order{
		CustomerId: tracker.CustomerId,
		Items:      tracker.Items,
	}

	return order, res.NewSuccessResponse("", http.StatusOK, s.probe.BaseProbe)
}

func (s *OrderService) GetLocalProbes() []probe.LocalProbe {
	return []probe.LocalProbe{*s.probe}
}

func (s *OrderService) GetDependencyProbes() []probe.DependencyProbe {
	return []probe.DependencyProbe{*s.inventorySvc.GetProbe()}
}
