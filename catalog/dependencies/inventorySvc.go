package dependencies

import (
	"encoding/json"
	"fmt"
	mod "health-probe/models"
	"health-probe/probe"
	res "health-probe/response"
	"math/rand"
	"net/http"
)

type InventorySvc struct {
	capacity        int
	inventorySvcUrl string
	probe           *probe.DependencyProbe
}

const getItem = "/items/{%d}"

func NewInventorySvc(capacity int, inventorySvcUrl string) *InventorySvc {
	probe := probe.NewDependencyProbe("CatalogSvc", "InventorySvc")
	return &InventorySvc{capacity: capacity, inventorySvcUrl: inventorySvcUrl, probe: probe}
}

func (i *InventorySvc) GetRandomItem() (mod.Item, res.ServiceResponse) {
	itemId := i.randomInRange()

	url := fmt.Sprintf(i.inventorySvcUrl+getItem, itemId)

	resp, err := http.Get(url)

	if err != nil {
		return mod.Item{}, res.NewErrorResponse(fmt.Sprintf("Error when getting suggestion: %s", err.Error()),
			http.StatusInternalServerError, i.probe.BaseProbe)
	}

	if resp.StatusCode != http.StatusOK {
		return mod.Item{}, res.NewErrorResponse(fmt.Sprintf("Unexpected status when getting suggestion: %s", resp.Status),
			http.StatusInternalServerError, i.probe.BaseProbe)
	}

	var item mod.Item
	if err := json.NewDecoder(resp.Body).Decode(&item); err != nil {
		return mod.Item{}, res.NewErrorResponse(fmt.Sprintf("Error when decoding item: %s", err.Error()),
			http.StatusInternalServerError, i.probe.BaseProbe)
	}

	return item, res.NewSuccessResponse("", resp.StatusCode, i.probe.BaseProbe)
}

func (i *InventorySvc) GetProbe() *probe.DependencyProbe {
	return i.probe
}

func (i *InventorySvc) randomInRange() int {
	return rand.Intn(i.capacity-1) + 1
}
