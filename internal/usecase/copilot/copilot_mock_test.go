package copilot

import (
	"hackbar-copilot/internal/domain/menu"
	"hackbar-copilot/internal/domain/recipe"
	"hackbar-copilot/internal/domain/stock"
	"hackbar-copilot/internal/usecase/sort"

	"github.com/stretchr/testify/mock"
)

var _ Copilot = new(MockCopilot)

type MockCopilot struct {
	mock.Mock
}

// FindGlassType implements Copilot.
func (m *MockCopilot) FindGlassType() (map[string]recipe.GlassType, error) {
	args := m.Called()
	return args.Get(0).(map[string]recipe.GlassType), args.Error(1)
}

// FindRecipeGroup implements Copilot.
func (m *MockCopilot) FindRecipeGroup(name string) (recipe.RecipeGroup, error) {
	args := m.Called(name)
	return args.Get(0).(recipe.RecipeGroup), args.Error(1)
}

// FindRecipeType implements Copilot.
func (m *MockCopilot) FindRecipeType() (map[string]recipe.RecipeType, error) {
	args := m.Called()
	return args.Get(0).(map[string]recipe.RecipeType), args.Error(1)
}

// ListMenu implements Copilot.
func (m *MockCopilot) ListMenu(sortFunc sort.Yield[menu.Group]) ([]menu.Group, error) {
	args := m.Called()
	return args.Get(0).([]menu.Group), args.Error(1)
}

// ListRecipes implements Copilot.
func (m *MockCopilot) ListRecipes(sortFunc sort.Yield[recipe.RecipeGroup]) ([]recipe.RecipeGroup, error) {
	args := m.Called()
	return args.Get(0).([]recipe.RecipeGroup), args.Error(1)
}

// SaveAsMenuGroup implements Copilot.
func (m *MockCopilot) SaveAsMenuGroup(recipeGroupName string, arg SaveAsMenuGroupArg) (menu.Group, error) {
	args := m.Called(recipeGroupName, arg)
	return args.Get(0).(menu.Group), args.Error(1)
}

// SaveGlassType implements Copilot.
func (m *MockCopilot) SaveGlassType(gt recipe.GlassType) error {
	args := m.Called(gt)
	return args.Error(0)
}

// SaveRecipe implements Copilot.
func (m *MockCopilot) SaveRecipe(rg recipe.RecipeGroup) error {
	args := m.Called(rg)
	return args.Error(0)
}

// SaveRecipeType implements Copilot.
func (m *MockCopilot) SaveRecipeType(rt recipe.RecipeType) error {
	args := m.Called(rt)
	return args.Error(0)
}

// Materials implements Copilot.
func (m *MockCopilot) Materials(sortFunc sort.Yield[stock.Material], optionAppliers ...QueryOptionApplier) ([]stock.Material, error) {
	args := m.Called()
	return args.Get(0).([]stock.Material), args.Error(1)
}

// UpdateStock implements Copilot.
func (m *MockCopilot) UpdateStock(inStockMaterials []string, outOfStockMaterials []string) error {
	args := m.Called(inStockMaterials, outOfStockMaterials)
	return args.Error(0)
}
