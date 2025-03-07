package recipeadapter

import (
	"hackbar-copilot/internal/domain/recipe"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	"hackbar-copilot/internal/usecase/copilot"
)

type InputAdapter interface {
	ApplyRecipeGroup(base recipe.RecipeGroup, input model.InputRecipeGroup) recipe.RecipeGroup
	ApplyRecipeTypes(
		current map[string]recipe.RecipeType, input model.InputRecipeGroup,
	) ([]recipe.RecipeType, error)
	ApplyGlassTypes(
		current map[string]recipe.GlassType, input model.InputRecipeGroup,
	) ([]recipe.GlassType, error)
	ApplyAsMenu(
		input model.InputRecipeGroup,
	) (asMenuArg *copilot.SaveAsMenuItemArg, err error)
}

func NewInputAdapter() InputAdapter {
	return &inputAdapter{}
}

type inputAdapter struct{}
