package inventory_service

import (
	cmn "health-probe/common"
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

func (inventorySvc *InventoryService) GetItem(id string) (Item, cmn.ServiceResponse) {
	inventorySvc.mu.Lock()
	defer inventorySvc.mu.Unlock()

	item, exists := inventorySvc.inventory[id]
	if !exists {
		return Item{}, cmn.NewErrorResponse("item not found", http.StatusNotFound)
	}

	return item, cmn.NewSuccessResponse("", http.StatusOK)
}

func (inventorySvc *InventoryService) AddItem(item Item) cmn.ServiceResponse {
	inventorySvc.mu.Lock()
	defer inventorySvc.mu.Unlock()

	_, exists := inventorySvc.inventory[item.ItemId]

	if exists {
		return cmn.NewErrorResponse("item already exists", http.StatusBadRequest)
	}

	inventorySvc.inventory[item.ItemId] = item
	return cmn.NewSuccessResponse("", http.StatusCreated)
}

func (inventorySvc *InventoryService) UpdateItem(item Item) cmn.ServiceResponse {
	inventorySvc.mu.Lock()
	defer inventorySvc.mu.Unlock()

	_, exists := inventorySvc.inventory[item.ItemId]

	if !exists {
		return cmn.NewErrorResponse("item not found", http.StatusNotFound)
	}

	inventorySvc.inventory[item.ItemId] = item
	return cmn.NewSuccessResponse("", http.StatusCreated)
}

func (inventorySvc *InventoryService) DeleteItem(id string) cmn.ServiceResponse {
	inventorySvc.mu.Lock()
	defer inventorySvc.mu.Unlock()

	_, exists := inventorySvc.inventory[id]

	if !exists {
		return cmn.NewErrorResponse("item not found", http.StatusNotFound)
	}

	delete(inventorySvc.inventory, id)
	return cmn.NewSuccessResponse("", http.StatusCreated)
}
