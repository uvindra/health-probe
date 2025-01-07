package order_service

import (
	res "health-probe/response"
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

func (s *OrderService) placeOrder(order Order) res.ServiceResponse {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.orders[order.OrderId]

	if exists {
		return res.NewErrorResponse("order already exists", http.StatusBadRequest)
	}

	s.orders[order.OrderId] = order
	return res.NewSuccessResponse("", http.StatusCreated)
}

func (s *OrderService) fetchOrder(id string) (Order, res.ServiceResponse) {
	s.mu.Lock()
	defer s.mu.Unlock()

	order, exists := s.orders[id]
	if !exists {
		return Order{}, res.NewErrorResponse("order not found", http.StatusNotFound)
	}

	return order, res.NewSuccessResponse("", http.StatusOK)
}
