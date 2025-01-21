package store

import (
	gen "health-probe/generator"
	mod "health-probe/models"
	"log"
	"sync"
)

type InventoryStore interface {
	GetItem(id int) (mod.Item, bool)
	GetItems() []mod.Item
	PopulateData()
	DeductItemQty(id int, qty int) bool
}

type inventoryStore struct {
	mu        sync.Mutex
	items     []mod.Item
	capacity  int
	reserve   int
	populated bool
}

func NewInventoryStore(capacity int, reserve int) *inventoryStore {
	return &inventoryStore{
		items:     make([]mod.Item, capacity+1), // +1 to capacity to avoid index out of range since ids will start from 1
		capacity:  capacity,
		reserve:   reserve,
		populated: false,
	}
}

func (d *inventoryStore) GetItem(id int) (mod.Item, bool) {
	d.validate(id)

	d.mu.Lock()
	defer d.mu.Unlock()

	return d.items[id], true
}

func (d *inventoryStore) GetItems() []mod.Item {
	if !d.populated {
		log.Panic("data not populated")
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	return d.items[1:]
}

func (d *inventoryStore) PopulateData() {
	d.mu.Lock()
	defer d.mu.Unlock()

	// Start from 1 since 0th index is not used
	for i := 1; i < len(d.items); i++ {
		d.items[i] = mod.Item{
			Id:       i,
			Name:     gen.GetRandomName(),
			Quantity: d.reserve,
		}
	}

	d.populated = true
}

func (d *inventoryStore) DeductItemQty(id int, qty int) bool {
	d.validate(id)

	d.mu.Lock()
	defer d.mu.Unlock()

	if d.items[id].Quantity < qty {
		return false
	}

	d.items[id].Quantity -= qty
	return true
}

func (d *inventoryStore) validate(id int) {
	if !d.populated {
		log.Panic("data not populated")
	}

	if id == 0 || id > len(d.items) {
		log.Panic("invalid item id")
	}
}
