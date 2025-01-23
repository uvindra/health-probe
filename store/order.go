package store

import (
	"fmt"

	enum "health-probe/enum"

	mod "health-probe/models"
	"sync"
)

type tracker map[string]mod.OrderTracker

type orderStore struct {
	mu       sync.Mutex
	trackers map[string]tracker
	seeds    map[string]int
}

func NewOrderStore() *orderStore {
	return &orderStore{
		trackers: make(map[string]tracker),
		seeds:    make(map[string]int),
	}
}

func (s *orderStore) AddOrderTracker(order mod.Order, state enum.OrderState) {
	s.mu.Lock()
	defer s.mu.Unlock()

	seed := s.seeds[order.CustomerId]
	seed++
	orderId := order.CustomerId + "-" + fmt.Sprint(seed)

	trk := s.trackers[order.CustomerId]

	trk[orderId] = mod.OrderTracker{
		Id:         orderId,
		CustomerId: order.CustomerId,
		Items:      order.Items,
		State:      state,
	}

	s.seeds[order.CustomerId] = seed
}

func (s *orderStore) GetOrderTracker(customerId string, orderId string) (mod.OrderTracker, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	trk, ok := s.trackers[customerId]

	if !ok {
		return mod.OrderTracker{}, false
	}

	order, ok := trk[orderId]

	return order, ok
}

func (s *orderStore) GetOrderTrackers(customerId string) []mod.OrderTracker {
	s.mu.Lock()
	defer s.mu.Unlock()

	trk, ok := s.trackers[customerId]

	if !ok {
		return []mod.OrderTracker{}
	}

	orders := make([]mod.OrderTracker, 0, len(trk))

	for _, order := range trk {
		orders = append(orders, order)
	}

	return orders
}
