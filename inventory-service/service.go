package inventory_service

import (
	res "health-probe/response"
	"net/http"
	"sync"
)

type InventoryService struct {
	mu        sync.Mutex
	inventory map[string]Item
}

func NewInventoryService() *InventoryService {
	return &InventoryService{
		inventory: make(map[string]Item),
	}
}

func (inventorySvc *InventoryService) getItem(id string) (Item, res.ServiceResponse) {
	inventorySvc.mu.Lock()
	defer inventorySvc.mu.Unlock()

	item, exists := inventorySvc.inventory[id]
	if !exists {
		return Item{}, res.NewErrorResponse("item not found", http.StatusNotFound)
	}

	return item, res.NewSuccessResponse("", http.StatusOK)
}

func (inventorySvc *InventoryService) addItem(item Item) res.ServiceResponse {
	inventorySvc.mu.Lock()
	defer inventorySvc.mu.Unlock()

	_, exists := inventorySvc.inventory[item.ItemId]

	if exists {
		return res.NewErrorResponse("item already exists", http.StatusBadRequest)
	}

	inventorySvc.inventory[item.ItemId] = item
	return res.NewSuccessResponse("", http.StatusCreated)
}

func (inventorySvc *InventoryService) updateItem(item Item) res.ServiceResponse {
	inventorySvc.mu.Lock()
	defer inventorySvc.mu.Unlock()

	_, exists := inventorySvc.inventory[item.ItemId]

	if !exists {
		return res.NewErrorResponse("item not found", http.StatusNotFound)
	}

	inventorySvc.inventory[item.ItemId] = item
	return res.NewSuccessResponse("", http.StatusCreated)
}

func (inventorySvc *InventoryService) deleteItem(id string) res.ServiceResponse {
	inventorySvc.mu.Lock()
	defer inventorySvc.mu.Unlock()

	_, exists := inventorySvc.inventory[id]

	if !exists {
		return res.NewErrorResponse("item not found", http.StatusNotFound)
	}

	delete(inventorySvc.inventory, id)
	return res.NewSuccessResponse("", http.StatusCreated)
}
