package copilot

import (
	"fmt"
	"hackbar-copilot/internal/domain/stock"
	"hackbar-copilot/internal/usecase/sort"

	"github.com/tingtt/options"
)

// Materials implements Copilot.
func (c *copilot) Materials(sortFunc sort.Yield[stock.Material], optionAppliers ...QueryOptionApplier) ([]stock.Material, error) {
	option := options.Create(optionAppliers...)

	var root *sort.Node[stock.Material]
	for material, err := range c.datasource.Stock().All() {
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve materials: %w", err)
		}
		if option.filterByName != nil {
			if material.Name != *option.filterByName {
				root = sort.Insert(root, material, sortFunc)
			}
		} else {
			root = sort.Insert(root, material, sortFunc)
		}
	}
	materials := []stock.Material{}
	sort.InOrderTraversal(root, &materials)
	return materials, nil
}
