package customer

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	mod "health-probe/models"
	"log"
	"net/http"
	"time"
)

type Customer struct {
	customerId      string
	catalogSvcUrl   string
	orderSvcUrl     string
	itemsPerOrder   int
	quantityPerItem int
}

func NewCustomer(catalogSvcUrl string, orderSvcUrl string, itemsPerOrder int, quantityPerItem int) *Customer {
	return &Customer{
		customerId:      genCustomerId(),
		catalogSvcUrl:   catalogSvcUrl,
		orderSvcUrl:     orderSvcUrl,
		itemsPerOrder:   itemsPerOrder,
		quantityPerItem: quantityPerItem,
	}
}

func (c *Customer) BeginShopping(stop chan bool) {
	for {
		select {
		case <-stop:
			fmt.Println("Ending shopping for customer ", c.customerId)
			return
		default:
			c.beginOrdering()
			time.Sleep(2 * time.Second)
		}
	}
}

func (c *Customer) beginOrdering() {
	newOrder := mod.Order{}
	newOrder.CustomerId = c.customerId

	for i := 0; i < c.itemsPerOrder; i++ {
		item, err := c.getSuggestion()

		if err != nil {
			log.Printf("Error when getting suggestion: %s", err.Error())
			return
		}

		if item.Quantity > 0 {
			newItem := item
			newItem.Quantity = c.quantityPerItem
			newOrder.Items = append(newOrder.Items, newItem)
		} else {
			log.Printf("No more items to order")
		}
	}

	if len(newOrder.Items) > 0 {
		if err := c.placeOrder(newOrder); err != nil {
			log.Printf("Error when placing order: %s", err.Error())
		}
	}

}

func (c *Customer) getSuggestion() (mod.Item, error) {
	resp, err := http.Get(c.catalogSvcUrl + "/suggest")
	if err != nil {
		return mod.Item{}, err
	}
	defer resp.Body.Close()

	var item mod.Item

	if err := json.NewDecoder(resp.Body).Decode(&item); err != nil {
		return mod.Item{}, err
	}

	return item, nil
}

func (c *Customer) placeOrder(order mod.Order) error {
	payload, err := json.Marshal(order)

	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, c.orderSvcUrl+"/order", bytes.NewBuffer(payload))

	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status when placing order: %s", resp.Status)
	} else {
		return nil
	}
}

func genCustomerId() string {
	b := make([]byte, 4) //equals 8 characters
	rand.Read(b)
	return hex.EncodeToString(b)
}
