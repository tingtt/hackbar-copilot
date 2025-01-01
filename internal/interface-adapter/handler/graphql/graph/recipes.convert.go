package graph

import (
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	"hackbar-copilot/internal/usecase/recipes"
	"hackbar-copilot/internal/utils/sliceutil"
)

type converterI interface {
	RecipeGroups(
		recipeGroups []recipes.RecipeGroup,
		recipeTypes map[string]model.RecipeType,
		glassTypes map[string]model.GlassType,
	) []*model.RecipeGroup
	RecipeGroup(
		recipeTypes map[string]model.RecipeType,
		glassTypes map[string]model.GlassType,
	) func(recipes.RecipeGroup) *model.RecipeGroup
	recipe(
		recipeTypes map[string]model.RecipeType,
		glassTypes map[string]model.GlassType,
	) func(recipe recipes.Recipe) *model.Recipe
}

var _ converterI = &converter{}

type converter struct{}

// RecipeGroups implements converterI.
func (convert *converter) RecipeGroups(recipeGroups []recipes.RecipeGroup, recipeTypes map[string]model.RecipeType, glassTypes map[string]model.GlassType) []*model.RecipeGroup {
	return sliceutil.Map(recipeGroups, convert.RecipeGroup(recipeTypes, glassTypes))
}

// RecipeGroup implements converterI.
func (convert *converter) RecipeGroup(recipeTypes map[string]model.RecipeType, glassTypes map[string]model.GlassType) func(recipes.RecipeGroup) *model.RecipeGroup {
	return func(recipeGroup recipes.RecipeGroup) *model.RecipeGroup {
		return &model.RecipeGroup{
			Name:     recipeGroup.Name,
			ImageURL: recipeGroup.ImageURL,
			Recipes:  sliceutil.Map(recipeGroup.Recipes, convert.recipe(recipeTypes, glassTypes)),
		}
	}
}

// recipe implements converterI.
func (convert *converter) recipe(recipeTypes map[string]model.RecipeType, glassTypes map[string]model.GlassType) func(recipe recipes.Recipe) *model.Recipe {
	return func(recipe recipes.Recipe) *model.Recipe {
		r := model.Recipe{
			Name:  recipe.Name,
			Steps: recipe.Steps,
		}
		recipeType, exists := recipeTypes[recipe.Type]
		if exists {
			r.Type = &recipeType
		}
		glassType, exists := glassTypes[recipe.Glass]
		if exists {
			r.Glass = &glassType
		}
		return &r
	}
}
