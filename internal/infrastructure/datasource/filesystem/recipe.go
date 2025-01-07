package filesystem

import (
	"hackbar-copilot/internal/domain/recipe"
	"iter"
)

var _ recipe.Repository = (*recipeRepository)(nil)

type recipeRepository struct {
	fs *filesystem
}

// All implements recipe.Repository.
func (r *recipeRepository) All() iter.Seq2[recipe.RecipeGroup, error] {
	return func(yield func(recipe.RecipeGroup, error) bool) {
		for _, rg := range r.fs.data.recipeGroups {
			if !yield(rg, nil) {
				return
			}
		}
	}
}

// AllGlassTypes implements recipe.Repository.
func (r *recipeRepository) AllGlassTypes() iter.Seq2[recipe.GlassType, error] {
	return func(yield func(recipe.GlassType, error) bool) {
		for _, gt := range r.fs.data.glassTypes {
			if !yield(gt, nil) {
				return
			}
		}
	}
}

// AllRecipeTypes implements recipe.Repository.
func (r *recipeRepository) AllRecipeTypes() iter.Seq2[recipe.RecipeType, error] {
	return func(yield func(recipe.RecipeType, error) bool) {
		for _, rt := range r.fs.data.recipeTypes {
			if !yield(rt, nil) {
				return
			}
		}
	}
}

// Save implements recipe.Repository.
func (r *recipeRepository) Save(new recipe.RecipeGroup) error {
	for i, savedRecipeGroup := range r.fs.data.recipeGroups {
		if savedRecipeGroup.Name == new.Name {
			r.fs.data.recipeGroups[i] = new
			return nil
		}
	}
	r.fs.data.recipeGroups = append(r.fs.data.recipeGroups, new)
	return nil
}

// SaveGlassType implements recipe.Repository.
func (r *recipeRepository) SaveGlassType(new recipe.GlassType) error {
	r.fs.data.glassTypes[new.Name] = new
	return nil
}

// SaveRecipeType implements recipe.Repository.
func (r *recipeRepository) SaveRecipeType(new recipe.RecipeType) error {
	r.fs.data.recipeTypes[new.Name] = new
	return nil
}
