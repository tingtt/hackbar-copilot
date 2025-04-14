package recipe

import (
	"iter"

	"github.com/stretchr/testify/mock"
)

var _ Repository = (*MockRepository)(nil)

type MockRepository struct {
	mock.Mock
}

// All implements Repository.
func (m *MockRepository) All() iter.Seq2[RecipeGroup, error] {
	args := m.Called()
	return args.Get(0).(iter.Seq2[RecipeGroup, error])
}

// AllGlassTypes implements Repository.
func (m *MockRepository) AllGlassTypes() iter.Seq2[GlassType, error] {
	args := m.Called()
	return args.Get(0).(iter.Seq2[GlassType, error])
}

// AllRecipeTypes implements Repository.
func (m *MockRepository) AllRecipeTypes() iter.Seq2[RecipeType, error] {
	args := m.Called()
	return args.Get(0).(iter.Seq2[RecipeType, error])
}

// Save implements Repository.
func (m *MockRepository) Save(rg RecipeGroup) error {
	args := m.Called(rg)
	return args.Error(0)
}

// SaveGlassType implements Repository.
func (m *MockRepository) SaveGlassType(gt GlassType) error {
	args := m.Called(gt)
	return args.Error(0)
}

// SaveRecipeType implements Repository.
func (m *MockRepository) SaveRecipeType(rt RecipeType) error {
	args := m.Called(rt)
	return args.Error(0)
}

// Remove implements Repository.
func (m *MockRepository) Remove(recipeGroupName string) error {
	args := m.Called(recipeGroupName)
	return args.Error(0)
}
