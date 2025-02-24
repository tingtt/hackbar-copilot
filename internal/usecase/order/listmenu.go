package order

import (
	"hackbar-copilot/internal/domain/menu"
	"hackbar-copilot/internal/usecase/sort"
)

// ListMenu implements Order.
func (o *orderimpl) ListMenu(sortFunc sort.Yield[menu.Group]) ([]menu.Group, error) {
	var root *sort.Node[menu.Group]
	for rg, err := range o.menu.All() {
		if err != nil {
			return nil, err
		}
		root = sort.Insert(root, rg, sortFunc)
	}
	menuGroups := []menu.Group{}
	sort.InOrderTraversal(root, &menuGroups)
	return menuGroups, nil
}
