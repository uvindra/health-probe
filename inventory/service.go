package inventory

import (
	mod "health-probe/models"
	"health-probe/probe"
	res "health-probe/response"
	store "health-probe/store"
	"net/http"
)

type InventoryService struct {
	inventory store.InventoryStore
	probe     *probe.LocalProbe
}

func NewService(capacity int, reserve int) *InventoryService {
	svc := &InventoryService{
		inventory: store.NewInventoryStore(capacity, reserve),
	}

	svc.inventory.PopulateData()
	return svc
}

func (i *InventoryService) getItems() ([]mod.Item, res.ServiceResponse) {
	items := i.inventory.GetItems()

	if len(items) == 0 {
		return nil, res.NewErrorResponse("Items do not exist", http.StatusNotFound, i.probe.BaseProbe)
	}

	return items, res.NewSuccessResponse("", http.StatusOK, i.probe.BaseProbe)
}

func (i *InventoryService) getItem(id int) (*mod.Item, res.ServiceResponse) {
	item, ok := i.inventory.GetItem(id)

	if !ok {
		return nil, res.NewErrorResponse("Item not found", http.StatusNotFound, i.probe.BaseProbe)
	}

	return item, res.NewSuccessResponse("", http.StatusOK, i.probe.BaseProbe)
}

func (i *InventoryService) deductItemQty(id int, quantity int) res.ServiceResponse {
	ok := i.inventory.DeductItemQty(id, quantity)

	if !ok {
		return res.NewErrorResponse("Item not found", http.StatusNotFound, i.probe.BaseProbe)
	}

	return res.NewSuccessResponse("", http.StatusOK, i.probe.BaseProbe)
}

func (i *InventoryService) GetLocalProbes() []probe.LocalProbe {
	return []probe.LocalProbe{*i.probe}
}

func (i *InventoryService) GetDependencyProbes() []probe.DependencyProbe {
	return nil
}
