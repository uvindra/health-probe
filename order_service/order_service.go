package order_service

import (
	cmn "health-probe/common"
	"net/http"
	"sync"
)

type OrderService struct {
	mu     sync.Mutex
	orders map[string]Order
}

func NewOrderService() *OrderService {
	return &OrderService{
		orders: make(map[string]Order),
	}
}

func (s *OrderService) PlaceOrder(order Order) *cmn.ServiceError {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.orders[order.OrderId]

	if exists {
		return cmn.NewServiceError("order already exists", http.StatusBadRequest)
	}

	s.orders[order.OrderId] = order
	return nil
}

func (s *OrderService) FetchOrder(id string) (Order, *cmn.ServiceError) {
	s.mu.Lock()
	defer s.mu.Unlock()

	order, exists := s.orders[id]
	if !exists {
		return Order{}, cmn.NewServiceError("order not found", http.StatusNotFound)
	}

	return order, nil
}
