package filesystem

import (
	"hackbar-copilot/internal/domain/recipe"
	"iter"
	"slices"
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
	if r.fs.data.glassTypes == nil {
		r.fs.data.glassTypes = make(map[string]recipe.GlassType)
	}
	r.fs.data.glassTypes[new.Name] = new
	return nil
}

// SaveRecipeType implements recipe.Repository.
func (r *recipeRepository) SaveRecipeType(new recipe.RecipeType) error {
	if r.fs.data.recipeTypes == nil {
		r.fs.data.recipeTypes = make(map[string]recipe.RecipeType)
	}
	r.fs.data.recipeTypes[new.Name] = new
	return nil
}

// Remove implements recipe.Repository.
func (r *recipeRepository) Remove(recipeGroupName string) error {
	for i, savedRecipeGroup := range r.fs.data.recipeGroups {
		if savedRecipeGroup.Name == recipeGroupName {
			r.fs.data.recipeGroups = slices.Delete(r.fs.data.recipeGroups, i, i+1)
			return nil
		}
	}
	return recipe.ErrNotFound
}
