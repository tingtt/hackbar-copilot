package copilot

import (
	"hackbar-copilot/internal/domain/menu"
	"hackbar-copilot/internal/domain/recipe"
	"hackbar-copilot/internal/domain/recipe/recipetest"
	"hackbar-copilot/internal/domain/stock/stocktest"
	usecaseutils "hackbar-copilot/internal/usecase/utils"
	"iter"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_copilot_SaveAsMenuItem(t *testing.T) {
	t.Parallel()

	t.Run("will call copilot.FindRecipeGroup with recipeGroupName", func(t *testing.T) {
		t.Parallel()

		t.Run("found", func(t *testing.T) {
			t.Parallel()

			recipe := new(MockRecipe)
			recipe.On("All").Return(recipetest.ExampleRecipeGroupsIter, nil)
			menu := new(MockMenu)
			menu.On("Save", mock.Anything).Return(nil)
			stock := new(MockStock)
			stock.On("All").Return(stocktest.ExampleMaterialsIter, nil)
			stock.On("Save", mock.Anything, mock.Anything).Return(nil)
			gateway := MockGateway{recipe: recipe, menu: menu, stock: stock}

			c := &copilot{&gateway}
			got, err := c.SaveAsMenuItem("Phuket Sling", SaveAsMenuItemArg{})
			assert.NoError(t, err)
			assert.Equal(t, "Phuket Sling", got.Name)
		})

		t.Run("not found", func(t *testing.T) {
			t.Parallel()

			recipe := new(MockRecipe)
			recipe.On("All").Return(recipetest.ExampleRecipeGroupsIter, nil)
			gateway := MockGateway{recipe: recipe}

			c := &copilot{&gateway}
			_, err := c.SaveAsMenuItem("-", SaveAsMenuItemArg{})
			assert.ErrorIs(t, err, usecaseutils.ErrNotFound)
		})
	})

	t.Run("will return menu group converted from recipe group", func(t *testing.T) {
		t.Parallel()

		recipe := new(MockRecipe)
		recipe.On("All").Return(recipetest.ExampleRecipeGroupsIter, nil)
		_menu := new(MockMenu)
		_menu.On("Save", mock.Anything).Return(nil)
		stock := new(MockStock)
		stock.On("All").Return(stocktest.ExampleMaterialsIter, nil)
		gateway := MockGateway{recipe: recipe, menu: _menu, stock: stock}

		c := &copilot{&gateway}
		got, err := c.SaveAsMenuItem("Phuket Sling", SaveAsMenuItemArg{
			Flavor: new("Sweet"),
			Options: map[string]MenuFromRecipeGroupArg{
				"Cocktail": {
					ImageURL: new("https://example.com/path/to/image/phuket-sling"),
					Price:    700,
				},
			},
		})
		want := menu.Item{
			Name:     "Phuket Sling",
			ImageURL: new("https://example.com/path/to/image/phuket-sling"),
			Flavor:   new("Sweet"),
			Options: []menu.ItemOption{
				{
					Name:       "Cocktail",
					Category:   "Cocktail",
					ImageURL:   new("https://example.com/path/to/image/phuket-sling"),
					Materials:  []string{"Peach liqueur", "Blue curacao", "Grapefruit juice", "Tonic water"},
					OutOfStock: false,
					Price:      700,
				},
			},
		}
		assert.NoError(t, err)
		assert.Equal(t, want, got)
		_menu.AssertCalled(t, "Save", want)
		stock.AssertNotCalled(t, "Save")
	})

	t.Run("will save new material", func(t *testing.T) {
		t.Parallel()

		_recipe := new(MockRecipe)
		recipeGroups := func() iter.Seq2[recipe.RecipeGroup, error] {
			return func(yield func(recipe.RecipeGroup, error) bool) {
				for rg := range recipetest.ExampleRecipeGroupsIter {
					if !yield(rg, nil) {
						return
					}
				}
				yield(recipe.RecipeGroup{
					Name:     "New Recipe",
					ImageURL: new("https://example.com/path/to/image/new"),
					Recipes: []recipe.Recipe{
						{
							Name:  "New Recipe",
							Type:  "build",
							Glass: "collins",
							Steps: []recipe.Step{
								{
									Material: new("New Material"),
									Amount:   new("30ml"),
								},
							},
						},
					},
				}, nil)
			}
		}()
		_recipe.On("All").Return(recipeGroups, nil)
		_menu := new(MockMenu)
		_menu.On("Save", mock.Anything).Return(nil)
		stock := new(MockStock)
		stock.On("All").Return(stocktest.ExampleMaterialsIter, nil)
		stock.On("Save", mock.Anything, mock.Anything).Return(nil)
		gateway := MockGateway{recipe: _recipe, menu: _menu, stock: stock}

		c := &copilot{&gateway}
		got, err := c.SaveAsMenuItem("New Recipe", SaveAsMenuItemArg{
			Flavor: new("Fruity"),
			Options: map[string]MenuFromRecipeGroupArg{
				"New Recipe": {
					ImageURL: new("https://example.com/path/to/image/new-recipe"),
					Price:    700,
				},
			},
		})
		want := menu.Item{
			Name:     "New Recipe",
			ImageURL: new("https://example.com/path/to/image/new"),
			Flavor:   new("Fruity"),
			Options: []menu.ItemOption{
				{
					Name:       "New Recipe",
					ImageURL:   new("https://example.com/path/to/image/new-recipe"),
					Materials:  []string{"New Material"},
					OutOfStock: false,
					Price:      700,
				},
			},
		}
		assert.NoError(t, err)
		assert.Equal(t, want, got)
		_menu.AssertCalled(t, "Save", want)
		stock.AssertCalled(t, "Save", []string{"New Material"}, mock.Anything)
	})

	t.Run("will call menu.Remove with recipeGroupName", func(t *testing.T) {
		t.Parallel()

		menuMock := new(MockMenu)
		menuMock.On("Remove", mock.Anything).Return(nil)
		gateway := MockGateway{menu: menuMock}

		c := &copilot{&gateway}
		_, err := c.SaveAsMenuItem("Phuket Sling", SaveAsMenuItemArg{Remove: true})
		assert.NoError(t, err)
		menuMock.AssertCalled(t, "Remove", "Phuket Sling")
	})
}
