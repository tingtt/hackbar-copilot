package copilot

import (
	"hackbar-copilot/internal/domain/menu"
	"hackbar-copilot/internal/usecase/sort"
)

func SortMenuGroupByName(fallbacks ...sort.YieldMaker[menu.Group]) sort.Yield[menu.Group] {
	return func(new, curr menu.Group) (isLeft bool) {
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
