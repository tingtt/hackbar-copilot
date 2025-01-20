package recipeadapter

import (
	"hackbar-copilot/internal/domain/recipe"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/adapter"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	"hackbar-copilot/internal/utils/sliceutil"
	"slices"
)

// ApplyGlassTypes implements Service.
func (s *inputAdapter) ApplyGlassTypes(
	current map[string]recipe.GlassType,
	input model.InputRecipeGroup,
) ([]recipe.GlassType, error) {
	return s.HasUpdateGlassTypes(
		s.ExtractGlassTypes(sliceutil.FilterNonNilPointerValues(input.Recipes)),
		current,
	)
}

func (s *inputAdapter) ExtractGlassTypes(recipes []model.InputRecipe) []model.InputGlassType {
	glassTypes := make([]model.InputGlassType, 0, len(recipes))
	for _, recipe := range recipes {
		if recipe.GlassType != nil {
			glassTypes = append(glassTypes, *recipe.GlassType)
		}
	}
	return glassTypes
}

func (s *inputAdapter) HasUpdateGlassTypes(inputGlassTypes []model.InputGlassType, currentGlassTypes map[string]recipe.GlassType) ([]recipe.GlassType, error) {
	hasUpdateGlassTypes := make([]recipe.GlassType, 0, len(inputGlassTypes))
	for _, inputGlassType := range inputGlassTypes {
		if /* GlassType already exists */ _, exists := currentGlassTypes[inputGlassType.Name]; exists {
			if /* has updates */ inputGlassType.ImageURL != nil || inputGlassType.Description != nil {
				if /* save is not true */ !(inputGlassType.Save != nil && *inputGlassType.Save) {
					return nil, adapter.ErrSaveDoesNotExist
				}
			} else {
				continue
			}
		}
		hasUpdateGlassTypes = append(hasUpdateGlassTypes, recipe.GlassType{
			Name:        inputGlassType.Name,
			ImageURL:    inputGlassType.ImageURL,
			Description: inputGlassType.Description,
		})
	}
	return slices.Clip(hasUpdateGlassTypes), nil
}
