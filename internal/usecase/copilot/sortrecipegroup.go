package copilot

import (
	"hackbar-copilot/internal/domain/recipe"
	"hackbar-copilot/internal/usecase/sort"
)

func SortRecipeGroupByName(fallbacks ...sort.YieldMaker[recipe.RecipeGroup]) sort.Yield[recipe.RecipeGroup] {
	return func(new, curr recipe.RecipeGroup) (isLeft bool) {
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
