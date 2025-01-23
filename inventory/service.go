package inventory

import (
	"fmt"
	mod "health-probe/models"
	res "health-probe/response"
	store "health-probe/store"
	"net/http"
)

type InventoryService struct {
	inventory store.InventoryStore
}

func NewService(capacity int, reserve int) *InventoryService {
	svc := &InventoryService{
		inventory: store.NewInventoryStore(capacity, reserve),
	}

	svc.inventory.PopulateData()
	return svc
}

func (inventorySvc *InventoryService) getItems() ([]mod.Item, res.ServiceResponse) {
	items := inventorySvc.inventory.GetItems()

	return items, res.NewSuccessResponse("", http.StatusOK)
}

func (inventorySvc *InventoryService) getItem(id int) (mod.Item, res.ServiceResponse) {
	item, exists := inventorySvc.inventory.GetItem(id)
	if !exists {
		return mod.Item{}, res.NewErrorResponse("item not found", http.StatusNotFound)
	}

	return item, res.NewSuccessResponse("", http.StatusOK)
}

func (inventorySvc *InventoryService) deductItemQty(id int, quantity int) res.ServiceResponse {
	ok := inventorySvc.inventory.DeductItemQty(id, quantity)

	if !ok {
		return res.NewErrorResponse(fmt.Sprintf("Dedcut qty:%d  is greater than current reserve", quantity),
			http.StatusBadRequest)
	}

	return res.NewSuccessResponse("", http.StatusOK)
}
