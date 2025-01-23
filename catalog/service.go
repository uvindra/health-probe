package catalog

import (
	"encoding/json"
	"fmt"
	mod "health-probe/models"
	res "health-probe/response"
	"math/rand"
	"net/http"
)

type Catalog struct {
	capacity        int
	inventorySvcUrl string
}

const getItem = "/items/{%d}"

func NewService(capacity int, inventorySvcUrl string) *Catalog {
	return &Catalog{capacity: capacity, inventorySvcUrl: inventorySvcUrl}
}

func (c *Catalog) GetSuggestion() (mod.Item, res.ServiceResponse) {
	itemId := c.randomInRange()

	url := fmt.Sprintf(c.inventorySvcUrl+getItem, itemId)

	resp, err := http.Get(url)

	if err != nil {
		return mod.Item{}, res.NewErrorResponse(fmt.Sprintf("Error when getting suggestion: %s", err.Error()),
			http.StatusInternalServerError)
	}

	if resp.StatusCode != http.StatusOK {
		return mod.Item{}, res.NewErrorResponse(fmt.Sprintf("Unexpected status when getting suggestion: %s", resp.Status),
			http.StatusInternalServerError)
	}

	var item mod.Item
	if err := json.NewDecoder(resp.Body).Decode(&item); err != nil {
		return mod.Item{}, res.NewErrorResponse(fmt.Sprintf("Error when decoding item: %s", err.Error()),
			http.StatusInternalServerError)
	}

	return item, res.NewSuccessResponse("", resp.StatusCode)
}

func (c *Catalog) randomInRange() int {
	return rand.Intn(c.capacity-1) + 1
}
