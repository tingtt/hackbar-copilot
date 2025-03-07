package recipeadapter

import (
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	"hackbar-copilot/internal/usecase/copilot"
)

// ApplyAsMenu implements RecipeAdapterIn.
func (s *inputAdapter) ApplyAsMenu(input model.InputRecipeGroup) (asMenuArg *copilot.SaveAsMenuItemArg, err error) {
	asMenuArg = &copilot.SaveAsMenuItemArg{}
	if input.AsMenu != nil {
		asMenuArg.Flavor = input.AsMenu.Flavor
	}
	items := make(map[string]copilot.MenuFromRecipeGroupArg)
	for _, recipe := range input.Recipes {
		if recipe.AsMenu != nil {
			items[recipe.Name] = copilot.MenuFromRecipeGroupArg{
				ImageURL: recipe.AsMenu.ImageURL,
				Price:    float32(recipe.AsMenu.Price),
			}
		}
	}
	asMenuArg.Options = items
	return asMenuArg, nil
}
