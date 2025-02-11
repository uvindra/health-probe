package catalog

import (
	intern "health-probe/catalog/dependencies"
	mod "health-probe/models"
	probe "health-probe/probe"
	res "health-probe/response"
)

type Catalog struct {
	inventorySvc *intern.InventorySvc
}

func NewService(capacity int, inventorySvcUrl string) *Catalog {
	inventorySvc := intern.NewInventorySvc(capacity, inventorySvcUrl)
	return &Catalog{inventorySvc: inventorySvc}
}

func (c *Catalog) GetSuggestion() (mod.Item, res.ServiceResponse) {
	return c.inventorySvc.GetRandomItem()
}

func (c *Catalog) GetLocalProbes() []probe.LocalProbe {
	return nil
}

func (c *Catalog) GetDependencyProbes() []probe.DependencyProbe {
	return []probe.DependencyProbe{*c.inventorySvc.GetProbe()}
}
