package adapter

import (
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	"hackbar-copilot/internal/usecase/copilot"
)

// ApplyAsMenu implements RecipeAdapterIn.
func (s *recipeAdapterIn) ApplyAsMenu(input model.InputRecipeGroup) (asMenuArg *copilot.SaveAsMenuGroupArg, err error) {
	asMenuArg = &copilot.SaveAsMenuGroupArg{}
	if input.AsMenu != nil {
		asMenuArg.Flavor = input.AsMenu.Flavor
	}
	items := make(map[string]copilot.MenuFromRecipeGroupArg)
	for _, recipe := range input.Recipes {
		if recipe.AsMenu != nil {
			items[recipe.Name] = copilot.MenuFromRecipeGroupArg{
				ImageURL: recipe.AsMenu.ImageURL,
				Price:    recipe.AsMenu.Price,
			}
		}
	}
	asMenuArg.Items = items
	return asMenuArg, nil
}
