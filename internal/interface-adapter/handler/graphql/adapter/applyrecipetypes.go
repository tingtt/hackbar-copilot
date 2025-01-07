package adapter

import (
	"errors"
	"hackbar-copilot/internal/domain/recipe"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	"hackbar-copilot/internal/utils/sliceutil"
	"slices"
)

// ApplyRecipeTypes implements Service.
func (s *recipeAdapterIn) ApplyRecipeTypes(
	current map[string]recipe.RecipeType,
	input model.InputRecipeGroup,
) ([]recipe.RecipeType, error) {
	return s.HasUpdateRecipeTypes(
		s.ExtractRecipeTypes(sliceutil.FilterNonNilPointerValues(input.Recipes)),
		current,
	)
}

func (s *recipeAdapterIn) ExtractRecipeTypes(recipes []model.InputRecipe) []model.InputRecipeType {
	recipeTypes := make([]model.InputRecipeType, 0, len(recipes))
	for _, recipe := range recipes {
		if recipe.RecipeType != nil {
			recipeTypes = append(recipeTypes, *recipe.RecipeType)
		}
	}
	return recipeTypes
}

func (s *recipeAdapterIn) HasUpdateRecipeTypes(inputRecipeTypes []model.InputRecipeType, currentRecipeTypes map[string]recipe.RecipeType) ([]recipe.RecipeType, error) {
	hasUpdateRecipeTypes := make([]recipe.RecipeType, 0, len(inputRecipeTypes))
	for _, inputRecipeType := range inputRecipeTypes {
		if /* recipeType already exists */ _, exists := currentRecipeTypes[inputRecipeType.Name]; exists {
			if /* has updates */ inputRecipeType.Description != nil {
				if /* save is not true */ !(inputRecipeType.Save != nil && *inputRecipeType.Save) {
					return nil, errors.New("save does not true")
				}
			} else {
				continue
			}
		}
		hasUpdateRecipeTypes = append(hasUpdateRecipeTypes, recipe.RecipeType{
			Name:        inputRecipeType.Name,
			Description: inputRecipeType.Description,
		})
	}
	return slices.Clip(hasUpdateRecipeTypes), nil
}
