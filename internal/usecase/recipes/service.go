package recipes

import (
	"errors"
	"fmt"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	usecaseutils "hackbar-copilot/internal/usecase/utils"
	"iter"
	"maps"
	"slices"
)

func NewService(repository Repository) (Service, error) {
	return &service{repository}, nil
}

type service struct {
	repository Repository
}

// Find implements Service.
func (s *service) Find() ([]RecipeGroup, error) {
	return s.repository.Find()
}

// FindRecipeType implements Service.
func (s *service) FindRecipeType() (map[string]model.RecipeType, error) {
	return s.repository.FindRecipeType()
}

// FindGlassType implements Service.
func (s *service) FindGlassType() (map[string]model.GlassType, error) {
	return s.repository.FindGlassType()
}

// Register implements Service.
func (s *service) Register(input model.InputRecipeGroup) (RecipeGroup, error) {
	recipeGroup, err := s.repository.FindOne(input.Name)
	if err != nil && !errors.Is(err, usecaseutils.ErrNotFound) {
		return RecipeGroup{}, err
	}

	recipeGroup = mergeRecipeGroup(recipeGroup, input)

	inputRecipeTypes, inputGlassTypes := extractInputRecipeTypeAndGlassType(input)
	err = s.SaveRecipeTypes(inputRecipeTypes)
	if err != nil {
		return RecipeGroup{}, fmt.Errorf("recipe type: %w", err)
	}
	err = s.SaveGlassTypes(inputGlassTypes)
	if err != nil {
		return RecipeGroup{}, fmt.Errorf("glass type: %w", err)
	}

	err = s.repository.Save(recipeGroup)
	if err != nil {
		return RecipeGroup{}, err
	}
	return recipeGroup, nil
}

func (s *service) SaveRecipeTypes(inputRecipeTypes []model.InputRecipeType) error {
	for _, inputRecipeType := range inputRecipeTypes {
		rts, err := s.repository.FindRecipeType()
		if err != nil {
			return err
		}
		if /* recipeType already exists */ _, exists := rts[inputRecipeType.Name]; exists {
			if inputRecipeType.Description != nil {
				if /* save is not true */ !(inputRecipeType.Save != nil && *inputRecipeType.Save) {
					return errors.New("save does not true")
					// TODO: return error ?
				}
			}
		}
		err = s.repository.SaveRecipeType(model.RecipeType{
			Name:        inputRecipeType.Name,
			Description: inputRecipeType.Description,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *service) SaveGlassTypes(inputGlassTypes []model.InputGlassType) error {
	for _, inputGlassType := range inputGlassTypes {
		gts, err := s.repository.FindGlassType()
		if err != nil {
			return err
		}
		if /* glassType already exists */ _, exists := gts[inputGlassType.Name]; exists {
			if inputGlassType.ImageURL != nil || inputGlassType.Description != nil {
				if /* save is not true */ !(inputGlassType.Save != nil && *inputGlassType.Save) {
					return errors.New("save does not true")
					// TODO: return error ?
				}
			}
		}
		err = s.repository.SaveGlassType(model.GlassType{
			Name:        inputGlassType.Name,
			ImageURL:    inputGlassType.ImageURL,
			Description: inputGlassType.Description,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func mergeRecipeGroup(base RecipeGroup, input model.InputRecipeGroup) RecipeGroup {
	base.Name = input.Name
	if input.ImageURL != nil {
		base.ImageURL = input.ImageURL
	}
	baseRecipesMap := make(map[string]Recipe, len(base.Recipes))
	for _, recipe := range base.Recipes {
		baseRecipesMap[recipe.Name] = recipe
	}
	for baseRecipe, inputRecipe := range iterInputRecipes(base.Recipes, input.Recipes) {
		baseRecipesMap[inputRecipe.Name] = mergeRecipe(baseRecipe, inputRecipe)
	}
	base.Recipes = slices.Collect(maps.Values(baseRecipesMap))
	return base
}

func extractInputRecipeTypeAndGlassType(input model.InputRecipeGroup) ([]model.InputRecipeType, []model.InputGlassType) {
	recipeTypes := make([]model.InputRecipeType, 0, len(input.Recipes))
	glassTypes := make([]model.InputGlassType, 0, len(input.Recipes))
	for _, inputRecipe := range input.Recipes {
		if inputRecipe.RecipeType != nil {
			recipeTypes = append(recipeTypes, *inputRecipe.RecipeType)
		}
		if inputRecipe.GlassType != nil {
			glassTypes = append(glassTypes, *inputRecipe.GlassType)
		}
	}
	return recipeTypes, glassTypes
}

func mergeRecipe(baseRecipe *Recipe, inputRecipe *model.InputRecipe) Recipe {
	if baseRecipe == nil {
		baseRecipe = &Recipe{
			Name: inputRecipe.Name,
		}
	}
	if inputRecipe.RecipeType != nil {
		baseRecipe.Type = inputRecipe.RecipeType.Name
	}
	if inputRecipe.GlassType != nil {
		baseRecipe.Glass = inputRecipe.GlassType.Name
	}
	if inputRecipe.Steps != nil {
		baseRecipe.Steps = inputRecipe.Steps
	}
	return *baseRecipe
}

func iterInputRecipes(base []Recipe, input []*model.InputRecipe) iter.Seq2[*Recipe, *model.InputRecipe] {
	return func(yield func(*Recipe, *model.InputRecipe) bool) {
		for _, inputRecipe := range input {
			found := false
			for _, baseRecipe := range base {
				if baseRecipe.Name == inputRecipe.Name {
					if /* break */ !yield(&baseRecipe, inputRecipe) {
						return
					}
					found = true
					break
				}
			}
			if !found {
				if /* break */ !yield(nil, inputRecipe) {
					return
				}
			}
		}
	}
}
