package copilot

import (
	"hackbar-copilot/internal/domain/menu"
	"hackbar-copilot/internal/usecase/sort"
)

func SortMenuGroupByName(fallbacks ...sort.YieldMaker[menu.Item]) sort.Yield[menu.Item] {
	return func(new, curr menu.Item) (isLeft bool) {
		if new.Name == "" {
			return curr.Name == ""
		}
		if curr.Name == "" {
			return new.Name != ""
		}
		if new.Name == curr.Name {
			if fallbacks == nil {
				return false
			}
			return fallbacks[0](fallbacks[1:]...)(new, curr)
		}
		return new.Name < curr.Name
	}
}
