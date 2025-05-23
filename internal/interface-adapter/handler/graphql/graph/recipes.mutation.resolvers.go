package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.64

import (
	"context"
	"errors"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	usecaseutils "hackbar-copilot/internal/usecase/utils"
)

// SaveRecipe is the resolver for the saveRecipe field.
func (r *mutationResolver) SaveRecipe(ctx context.Context, input model.InputRecipeGroup) (model.SaveRecipeResult, error) {
	_, err := r.authAdapter.GetEmail(ctx)
	if /* unauthorized */ err != nil {
		return nil, err
	}
	if !r.authAdapter.HasBartenderRole(ctx) {
		return nil, errors.New("forbidden")
	}

	if input.Remove != nil && *input.Remove {
		err := r.Copilot.RemoveRecipeAndMenuItem(input.Name)
		return &model.RemovedRecipeGroup{Name: input.Name}, err
	}

	currentRecipeGroup, err := r.Copilot.FindRecipeGroup(input.Name)
	if err != nil && !errors.Is(err, usecaseutils.ErrNotFound) {
		return nil, err
	}
	currentRecipeTypes, err := r.Copilot.FindRecipeType()
	if err != nil {
		return nil, err
	}
	currentGlassTypes, err := r.Copilot.FindGlassType()
	if err != nil {
		return nil, err
	}
	newRecipeGroup := r.recipeAdapter.ApplyRecipeGroup(currentRecipeGroup, input)
	if len(newRecipeGroup.Recipes) == 0 {
		err := r.Copilot.RemoveRecipeAndMenuItem(input.Name)
		return &model.RemovedRecipeGroup{Name: input.Name}, err
	}

	newRecipeTypes, err := r.recipeAdapter.ApplyRecipeTypes(currentRecipeTypes, input)
	if err != nil {
		return nil, err
	}
	newGlassTypes, err := r.recipeAdapter.ApplyGlassTypes(currentGlassTypes, input)
	if err != nil {
		return nil, err
	}
	asMenuArg, err := r.recipeAdapter.ApplyAsMenu(input)
	if err != nil {
		return nil, err
	}

	for _, newRecipeType := range newRecipeTypes {
		err := r.Copilot.SaveRecipeType(newRecipeType)
		if err != nil {
			return nil, err
		}
	}
	for _, newGlassType := range newGlassTypes {
		err := r.Copilot.SaveGlassType(newGlassType)
		if err != nil {
			return nil, err
		}
	}
	err = r.Copilot.SaveRecipe(newRecipeGroup)
	if err != nil {
		return nil, err
	}
	if asMenuArg != nil {
		_, err := r.Copilot.SaveAsMenuItem(newRecipeGroup.Name, *asMenuArg)
		if err != nil {
			return nil, err
		}
	}

	recipeTypes, err := r.Copilot.FindRecipeType()
	if err != nil {
		return nil, err
	}
	glassTypes, err := r.Copilot.FindGlassType()
	if err != nil {
		return nil, err
	}
	return r.recipeAdapter.RecipeGroup(recipeTypes, glassTypes)(newRecipeGroup), nil
}
