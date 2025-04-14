package copilot

import (
	"hackbar-copilot/internal/domain/recipe"
	"iter"

	"github.com/stretchr/testify/mock"
)

var _ recipe.SaveListRemover = new(MockRecipeSaveListRemover)

type MockRecipeSaveListRemover struct {
	mock.Mock
}

// All implements recipe.SaveLister.
func (m *MockRecipeSaveListRemover) All() iter.Seq2[recipe.RecipeGroup, error] {
	args := m.Called()
	return args.Get(0).(iter.Seq2[recipe.RecipeGroup, error])
}

// AllGlassTypes implements recipe.SaveLister.
func (m *MockRecipeSaveListRemover) AllGlassTypes() iter.Seq2[recipe.GlassType, error] {
	args := m.Called()
	return args.Get(0).(iter.Seq2[recipe.GlassType, error])
}

// AllRecipeTypes implements recipe.SaveLister.
func (m *MockRecipeSaveListRemover) AllRecipeTypes() iter.Seq2[recipe.RecipeType, error] {
	args := m.Called()
	return args.Get(0).(iter.Seq2[recipe.RecipeType, error])
}

// Save implements recipe.SaveLister.
func (m *MockRecipeSaveListRemover) Save(rg recipe.RecipeGroup) error {
	args := m.Called(rg)
	return args.Error(0)
}

// SaveGlassType implements recipe.SaveLister.
func (m *MockRecipeSaveListRemover) SaveGlassType(gt recipe.GlassType) error {
	args := m.Called(gt)
	return args.Error(0)
}

// SaveRecipeType implements recipe.SaveLister.
func (m *MockRecipeSaveListRemover) SaveRecipeType(rt recipe.RecipeType) error {
	args := m.Called(rt)
	return args.Error(0)
}

// Remove implements recipe.SaveListRemover.
func (m *MockRecipeSaveListRemover) Remove(recipeGroupName string) error {
	args := m.Called(recipeGroupName)
	return args.Error(0)
}
