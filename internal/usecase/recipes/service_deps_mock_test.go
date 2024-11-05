package recipes

import (
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"

	"github.com/stretchr/testify/mock"
)

var _ Service = new(MockRecipeService)

type MockRecipeService struct {
	mock.Mock
	find           func() ([]RecipeGroup, error)
	findGlassType  func() (map[string]model.GlassType, error)
	findRecipeType func() (map[string]model.RecipeType, error)
	register       func(input model.InputRecipeGroup) (RecipeGroup, error)
}

// Find implements Service.
func (m *MockRecipeService) Find() ([]RecipeGroup, error) {
	args := m.Called()
	if m.find != nil {
		return m.find()
	}
	return args.Get(0).([]RecipeGroup), args.Error(1)
}

// FindGlassType implements Service.
func (m *MockRecipeService) FindGlassType() (map[string]model.GlassType, error) {
	args := m.Called()
	if m.findGlassType != nil {
		return m.findGlassType()
	}
	return args.Get(0).(map[string]model.GlassType), args.Error(1)
}

// FindRecipeType implements Service.
func (m *MockRecipeService) FindRecipeType() (map[string]model.RecipeType, error) {
	args := m.Called()
	if m.findRecipeType != nil {
		return m.findRecipeType()
	}
	return args.Get(0).(map[string]model.RecipeType), args.Error(1)
}

// Register implements Service.
func (m *MockRecipeService) Register(input model.InputRecipeGroup) (RecipeGroup, error) {
	args := m.Called(input)
	if m.register != nil {
		return m.register(input)
	}
	return args.Get(0).(RecipeGroup), args.Error(1)
}

var _ Repository = new(MockRepository)

type MockRepository struct {
	mock.Mock
}

// Find implements Repository.
func (m *MockRepository) Find() ([]RecipeGroup, error) {
	args := m.Called()
	return args.Get(0).([]RecipeGroup), args.Error(1)
}

// FindGlassType implements Repository.
func (m *MockRepository) FindGlassType() (map[string]model.GlassType, error) {
	args := m.Called()
	return args.Get(0).(map[string]model.GlassType), args.Error(1)
}

// FindOne implements Repository.
func (m *MockRepository) FindOne(name string) (RecipeGroup, error) {
	args := m.Called(name)
	return args.Get(0).(RecipeGroup), args.Error(1)
}

// FindRecipeType implements Repository.
func (m *MockRepository) FindRecipeType() (map[string]model.RecipeType, error) {
	args := m.Called()
	return args.Get(0).(map[string]model.RecipeType), args.Error(1)
}

// Save implements Repository.
func (m *MockRepository) Save(rg RecipeGroup) error {
	args := m.Called(rg)
	return args.Error(0)
}

// SaveGlassType implements Repository.
func (m *MockRepository) SaveGlassType(gt model.GlassType) error {
	args := m.Called(gt)
	return args.Error(0)
}

// SaveRecipeType implements Repository.
func (m *MockRepository) SaveRecipeType(rt model.RecipeType) error {
	args := m.Called(rt)
	return args.Error(0)
}
