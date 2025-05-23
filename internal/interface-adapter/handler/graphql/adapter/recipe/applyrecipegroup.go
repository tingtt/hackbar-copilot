package recipeadapter

import (
	"hackbar-copilot/internal/domain/recipe"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	"hackbar-copilot/internal/utils/sliceutil"
	"iter"
	"maps"
	"slices"
)

// ApplyRecipeGroup implements Service.
func (s *inputAdapter) ApplyRecipeGroup(base recipe.RecipeGroup, input model.InputRecipeGroup) recipe.RecipeGroup {
	if input.Name != "" {
		base.Name = input.Name
	}
	if input.ImageURL != nil {
		base.ImageURL = input.ImageURL
	}
	baseRecipesMap := make(map[string]recipe.Recipe, len(base.Recipes))

	if input.Replace == nil || !*input.Replace {
		for _, recipe := range base.Recipes {
			baseRecipesMap[recipe.Name] = recipe
		}
	}
	for baseRecipe, inputRecipe := range s.iterInputRecipes(base.Recipes, sliceutil.FilterNonNilPointerValues(input.Recipes)) {
		if inputRecipe.Remove != nil && *inputRecipe.Remove {
			delete(baseRecipesMap, baseRecipe.Name)
			continue
		}
		baseRecipesMap[inputRecipe.Name] = s.applyRecipe(baseRecipe, inputRecipe)
	}
	base.Recipes = slices.Collect(maps.Values(baseRecipesMap))
	return base
}

func (s *inputAdapter) iterInputRecipes(base []recipe.Recipe, input []model.InputRecipe) iter.Seq2[recipe.Recipe, model.InputRecipe] {
	return func(yield func(recipe.Recipe, model.InputRecipe) bool) {
		for _, inputRecipe := range input {
			found := false
			for _, baseRecipe := range base {
				if baseRecipe.Name == inputRecipe.Name {
					if /* break */ !yield(baseRecipe, inputRecipe) {
						return
					}
					found = true
					break
				}
			}
			if !found {
				if /* break */ !yield(recipe.Recipe{}, inputRecipe) {
					return
				}
			}
		}
	}
}

func (s *inputAdapter) applyRecipe(baseRecipe recipe.Recipe, inputRecipe model.InputRecipe) recipe.Recipe {
	baseRecipe.Name = inputRecipe.Name
	baseRecipe.Category = inputRecipe.Category
	if inputRecipe.RecipeType != nil {
		baseRecipe.Type = inputRecipe.RecipeType.Name
	}
	if inputRecipe.GlassType != nil {
		baseRecipe.Glass = inputRecipe.GlassType.Name
	}
	if inputRecipe.Steps != nil {
		baseRecipe.Steps = sliceutil.Map(inputRecipe.Steps, s.step)
	}
	return baseRecipe
}

func (s *inputAdapter) step(step *model.InputStep) recipe.Step {
	return recipe.Step{
		Material:    step.Material,
		Amount:      step.Amount,
		Description: step.Description,
	}
}
