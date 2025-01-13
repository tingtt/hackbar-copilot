package copilot

import (
	"hackbar-copilot/internal/domain/menu"
	"hackbar-copilot/internal/usecase/sort"
)

// ListMenu implements Copilot.
func (c *copilot) ListMenu(sortFunc sort.Yield[menu.Group]) ([]menu.Group, error) {
	var root *sort.Node[menu.Group]
	for rg, err := range c.menu.All() {
		if err != nil {
			return nil, err
		}
		root = sort.Insert(root, rg, sortFunc)
	}
	menuGroups := []menu.Group{}
	sort.InOrderTraversal(root, &menuGroups)
	return menuGroups, nil
}
