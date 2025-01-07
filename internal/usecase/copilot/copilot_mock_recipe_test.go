package copilot

import (
	"hackbar-copilot/internal/domain/recipe"
	"iter"

	"github.com/stretchr/testify/mock"
)

var _ recipe.SaveLister = new(MockRecipeSaveLister)

type MockRecipeSaveLister struct {
	mock.Mock
}

// All implements recipe.SaveLister.
func (m *MockRecipeSaveLister) All() iter.Seq2[recipe.RecipeGroup, error] {
	args := m.Called()
	return args.Get(0).(iter.Seq2[recipe.RecipeGroup, error])
}

// AllGlassTypes implements recipe.SaveLister.
func (m *MockRecipeSaveLister) AllGlassTypes() iter.Seq2[recipe.GlassType, error] {
	args := m.Called()
	return args.Get(0).(iter.Seq2[recipe.GlassType, error])
}

// AllRecipeTypes implements recipe.SaveLister.
func (m *MockRecipeSaveLister) AllRecipeTypes() iter.Seq2[recipe.RecipeType, error] {
	args := m.Called()
	return args.Get(0).(iter.Seq2[recipe.RecipeType, error])
}

// Save implements recipe.SaveLister.
func (m *MockRecipeSaveLister) Save(rg recipe.RecipeGroup) error {
	args := m.Called(rg)
	return args.Error(0)
}

// SaveGlassType implements recipe.SaveLister.
func (m *MockRecipeSaveLister) SaveGlassType(gt recipe.GlassType) error {
	args := m.Called(gt)
	return args.Error(0)
}

// SaveRecipeType implements recipe.SaveLister.
func (m *MockRecipeSaveLister) SaveRecipeType(rt recipe.RecipeType) error {
	args := m.Called(rt)
	return args.Error(0)
}

var _ recipe.Repository = new(MockRepository)

type MockRepository struct {
	mock.Mock
}

// All implements recipe.Repository.
func (m *MockRepository) All() iter.Seq2[recipe.RecipeGroup, error] {
	args := m.Called()
	return args.Get(0).(iter.Seq2[recipe.RecipeGroup, error])
}

// AllGlassTypes implements recipe.Repository.
func (m *MockRepository) AllGlassTypes() iter.Seq2[recipe.GlassType, error] {
	args := m.Called()
	return args.Get(0).(iter.Seq2[recipe.GlassType, error])
}

// AllRecipeTypes implements recipe.Repository.
func (m *MockRepository) AllRecipeTypes() iter.Seq2[recipe.RecipeType, error] {
	args := m.Called()
	return args.Get(0).(iter.Seq2[recipe.RecipeType, error])
}

// Save implements recipe.Repository.
func (m *MockRepository) Save(rg recipe.RecipeGroup) error {
	args := m.Called(rg)
	return args.Error(0)
}

// SaveGlassType implements recipe.Repository.
func (m *MockRepository) SaveGlassType(gt recipe.GlassType) error {
	args := m.Called(gt)
	return args.Error(0)
}

// SaveRecipeType implements recipe.Repository.
func (m *MockRepository) SaveRecipeType(rt recipe.RecipeType) error {
	args := m.Called(rt)
	return args.Error(0)
}
