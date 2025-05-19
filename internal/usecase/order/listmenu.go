package order

import (
	"hackbar-copilot/internal/domain/menu"
	"hackbar-copilot/internal/usecase/sort"
	"slices"
)

// ListMenu implements Order.
func (o *orderimpl) ListMenu(sortFunc sort.Yield[menu.Item]) ([]menu.Item, error) {
	var root *sort.Node[menu.Item]
	for item, err := range o.datasource.Menu().All() {
		if err != nil {
			return nil, err
		}
		slices.SortFunc(item.Options, func(a, b menu.ItemOption) int {
			if a.Name == b.Name {
				return 0
			}
			if a.Name < b.Name {
				return -1
			}
			return 1
		})
		root = sort.Insert(root, item, sortFunc)
	}
	menuGroups := []menu.Item{}
	sort.InOrderTraversal(root, &menuGroups)
	return menuGroups, nil
}
