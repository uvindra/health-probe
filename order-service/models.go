package order_service

type Item struct {
	ItemId      string `json:"item_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

type Order struct {
	OrderId    string `json:"order_id"`
	CustomerId string `json:"customer_id"`
	Items      []Item `json:"items"`
}
