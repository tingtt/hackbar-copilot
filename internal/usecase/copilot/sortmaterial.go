package copilot

import (
	"hackbar-copilot/internal/domain/stock"
	"hackbar-copilot/internal/usecase/sort"
)

func SortMaterialByName(fallbacks ...sort.YieldMaker[stock.Material]) sort.Yield[stock.Material] {
	return func(new, curr stock.Material) (isLeft bool) {
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
