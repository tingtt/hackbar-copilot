package recipeadapter

import (
	"hackbar-copilot/internal/domain/recipe"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	"hackbar-copilot/internal/utils/sliceutil"
)

// RecipeGroup implements RecipeAdapterOut.
func (adapter *outputAdapter) RecipeGroup(recipeTypes map[string]recipe.RecipeType, glassTypes map[string]recipe.GlassType) func(recipe.RecipeGroup) *model.RecipeGroup {
	return func(recipeGroup recipe.RecipeGroup) *model.RecipeGroup {
		return &model.RecipeGroup{
			Name:     recipeGroup.Name,
			ImageURL: recipeGroup.ImageURL,
			Recipes:  sliceutil.Map(recipeGroup.Recipes, adapter.recipe(recipeTypes, glassTypes)),
		}
	}
}

func (adapter *outputAdapter) recipe(recipeTypes map[string]recipe.RecipeType, glassTypes map[string]recipe.GlassType) func(recipe recipe.Recipe) *model.Recipe {
	return func(_r recipe.Recipe) *model.Recipe {
		r := model.Recipe{
			Name:     _r.Name,
			Category: _r.Category,
			Steps:    sliceutil.Map(_r.Steps, adapter.step),
		}
		recipeType, exists := recipeTypes[_r.Type]
		if exists {
			r.Type = &model.RecipeType{
				Name:        recipeType.Name,
				Description: recipeType.Description,
			}
		}
		glassType, exists := glassTypes[_r.Glass]
		if exists {
			r.Glass = &model.GlassType{
				Name:        glassType.Name,
				ImageURL:    glassType.ImageURL,
				Description: glassType.Description,
			}
		}
		return &r
	}
}

func (adapter *outputAdapter) step(step recipe.Step) *model.Step {
	return &model.Step{
		Material:    step.Material,
		Amount:      step.Amount,
		Description: step.Description,
	}
}
