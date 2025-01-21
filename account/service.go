package account

import mod "health-probe/models"

type Account struct {
	items map[int]mod.Item
}

func NewAccount() *Account {
	return &Account{
		items: make(map[int]mod.Item),
	}
}

func (a *Account) AddItem(item mod.Item) {
	a.items[item.Id] = item
}

func (a *Account) GetItems() []mod.Item {
	items := make([]mod.Item, 0, len(a.items))
	for _, item := range a.items {
		items = append(items, item)
	}
	return items
}

func (a *Account) GetItem(id int) (mod.Item, bool) {
	item, ok := a.items[id]
	return item, ok
}

func (a *Account) RemoveItem(id int) bool {
	if _, ok := a.items[id]; !ok {
		return false
	}
	delete(a.items, id)
	return true
}
