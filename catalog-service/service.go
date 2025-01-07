package catalog_service

import (
	"encoding/json"
	"fmt"
	res "health-probe/response"
	"net/http"
)

type CatalogService struct {
	categoryMap     map[string][]string
	inventorySvcUrl string
}

const getItem = "/items/%s"

func NewCatalogService(inventorySvcUrl string) *CatalogService {
	return &CatalogService{
		categoryMap:     make(map[string][]string),
		inventorySvcUrl: inventorySvcUrl,
	}
}

func (catalogSvc *CatalogService) getItemsByCategory(category string) ([]Item, res.ServiceResponse) {
	itemIds, exists := catalogSvc.categoryMap[category]
	if !exists {
		return nil, res.NewSuccessResponse("Category not found", http.StatusNotFound)
	}

	resource := catalogSvc.inventorySvcUrl + getItem

	var items []Item

	for _, itemId := range itemIds {
		url := fmt.Sprintf(resource, itemId)
		resp, err := http.Get(url)

		if err != nil {
			return nil, res.NewErrorResponse(fmt.Sprintf("Error when calling Inventory service: %s", err.Error()),
				http.StatusInternalServerError)
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, res.NewErrorResponse(fmt.Sprintf("Unexpected status when getting item from Inventory service: %s", resp.Status),
				http.StatusInternalServerError)
		}

		var item Item
		if err := json.NewDecoder(resp.Body).Decode(&item); err != nil {
			return nil, res.NewErrorResponse(fmt.Sprintf("Error reading response from Inventory service: %s", err.Error()),
				http.StatusInternalServerError)
		}

		items = append(items, item)
	}

	return items, res.NewSuccessResponse("", http.StatusOK)
}
