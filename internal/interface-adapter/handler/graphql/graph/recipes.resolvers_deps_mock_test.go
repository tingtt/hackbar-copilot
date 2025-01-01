package graph

import (
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	"hackbar-copilot/internal/usecase/recipes"

	"github.com/stretchr/testify/mock"
)

var _ recipes.Service = new(MockRecipeService)

type MockRecipeService struct {
	mock.Mock
}

// Find implements recipes.Service.
func (m *MockRecipeService) Find() ([]recipes.RecipeGroup, error) {
	args := m.Called()
	return args.Get(0).([]recipes.RecipeGroup), args.Error(1)
}

// FindGlassType implements recipes.Service.
func (m *MockRecipeService) FindGlassType() (map[string]model.GlassType, error) {
	args := m.Called()
	return args.Get(0).(map[string]model.GlassType), args.Error(1)
}

// FindRecipeType implements recipes.Service.
func (m *MockRecipeService) FindRecipeType() (map[string]model.RecipeType, error) {
	args := m.Called()
	return args.Get(0).(map[string]model.RecipeType), args.Error(1)
}

// Register implements recipes.Service.
func (m *MockRecipeService) Register(input model.InputRecipeGroup) (recipes.RecipeGroup, error) {
	args := m.Called(input)
	return args.Get(0).(recipes.RecipeGroup), args.Error(1)
}

var _ converterI = new(MockConverter)

type MockConverter struct {
	mock.Mock
}

// RecipeGroups implements converterI.
func (m *MockConverter) RecipeGroups(recipeGroups []recipes.RecipeGroup, recipeTypes map[string]model.RecipeType, glassTypes map[string]model.GlassType) []*model.RecipeGroup {
	args := m.Called(recipeGroups, recipeTypes, glassTypes)
	return args.Get(0).([]*model.RecipeGroup)
}

// recipe implements converterI.
func (m *MockConverter) recipe(recipeTypes map[string]model.RecipeType, glassTypes map[string]model.GlassType) func(recipe recipes.Recipe) *model.Recipe {
	args := m.Called(recipeTypes, glassTypes)
	return args.Get(0).(func(recipe recipes.Recipe) *model.Recipe)
}

// RecipeGroup implements converterI.
func (m *MockConverter) RecipeGroup(recipeTypes map[string]model.RecipeType, glassTypes map[string]model.GlassType) func(recipes.RecipeGroup) *model.RecipeGroup {
	args := m.Called(recipeTypes, glassTypes)
	return args.Get(0).(func(recipes.RecipeGroup) *model.RecipeGroup)
}

type MockConverterRecipeGroupFunc struct {
	mock.Mock
}

func (m *MockConverterRecipeGroupFunc) Run(recipeGroup recipes.RecipeGroup) *model.RecipeGroup {
	args := m.Called(recipeGroup)
	return args.Get(0).(*model.RecipeGroup)
}
