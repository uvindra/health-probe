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

func (s *OrderService) placeOrder(order Order) cmn.ServiceResponse {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.orders[order.OrderId]

	if exists {
		return cmn.NewErrorResponse("order already exists", http.StatusBadRequest)
	}

	s.orders[order.OrderId] = order
	return cmn.NewSuccessResponse("", http.StatusCreated)
}

func (s *OrderService) fetchOrder(id string) (Order, cmn.ServiceResponse) {
	s.mu.Lock()
	defer s.mu.Unlock()

	order, exists := s.orders[id]
	if !exists {
		return Order{}, cmn.NewErrorResponse("order not found", http.StatusNotFound)
	}

	return order, cmn.NewSuccessResponse("", http.StatusOK)
}
