package catalog_service

type Item struct {
	ItemId      string `json:"item_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
