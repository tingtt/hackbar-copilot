package recipetest

import (
	"hackbar-copilot/internal/domain/recipe"
	"iter"
)

func IterWithNilError(items []recipe.RecipeGroup) iter.Seq2[recipe.RecipeGroup, error] {
	return func(yield func(recipe.RecipeGroup, error) bool) {
		for _, item := range items {
			if !yield(item, nil) {
				break
			}
		}
	}
}

func IterWithNilErrorRecipeTypes(items []recipe.RecipeType) iter.Seq2[recipe.RecipeType, error] {
	return func(yield func(recipe.RecipeType, error) bool) {
		for _, item := range items {
			if !yield(item, nil) {
				break
			}
		}
	}
}

func IterWithNilErrorGlassTypes(items []recipe.GlassType) iter.Seq2[recipe.GlassType, error] {
	return func(yield func(recipe.GlassType, error) bool) {
		for _, item := range items {
			if !yield(item, nil) {
				break
			}
		}
	}
}
