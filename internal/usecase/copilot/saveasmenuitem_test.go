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

			recipeSaveLister := new(MockRecipeSaveListRemover)
			recipeSaveLister.On("All").Return(recipetest.ExampleRecipeGroupsIter, nil)

			menuSaveLister := new(MockMenu)
			menuSaveLister.On("Save", mock.Anything).Return(nil)

			stockSaveLister := new(MockStockSaveLister)
			stockSaveLister.On("All").Return(stocktest.ExampleMaterialsIter, nil)
			stockSaveLister.On("Save", mock.Anything, mock.Anything).Return(nil)

			c := &copilot{recipe: recipeSaveLister, menu: menuSaveLister, stock: stockSaveLister}
			got, err := c.SaveAsMenuItem("Phuket Sling", SaveAsMenuItemArg{})
			assert.NoError(t, err)
			assert.Equal(t, "Phuket Sling", got.Name)
		})

		t.Run("not found", func(t *testing.T) {
			t.Parallel()

			recipeSaveLister := new(MockRecipeSaveListRemover)
			recipeSaveLister.On("All").Return(recipetest.ExampleRecipeGroupsIter, nil)

			c := &copilot{recipe: recipeSaveLister}
			_, err := c.SaveAsMenuItem("-", SaveAsMenuItemArg{})
			assert.ErrorIs(t, err, usecaseutils.ErrNotFound)
		})
	})

	t.Run("will return menu group converted from recipe group", func(t *testing.T) {
		t.Parallel()

		recipeSaveLister := new(MockRecipeSaveListRemover)
		recipeSaveLister.On("All").Return(recipetest.ExampleRecipeGroupsIter, nil)

		menuSaveLister := new(MockMenu)
		menuSaveLister.On("Save", mock.Anything).Return(nil)

		stockSaveLister := new(MockStockSaveLister)
		stockSaveLister.On("All").Return(stocktest.ExampleMaterialsIter, nil)

		c := &copilot{recipe: recipeSaveLister, menu: menuSaveLister, stock: stockSaveLister}
		got, err := c.SaveAsMenuItem("Phuket Sling", SaveAsMenuItemArg{
			Flavor: ptr("Sweet"),
			Options: map[string]MenuFromRecipeGroupArg{
				"Cocktail": {
					ImageURL: ptr("https://example.com/path/to/image/phuket-sling"),
					Price:    700,
				},
			},
		})
		want := menu.Item{
			Name:     "Phuket Sling",
			ImageURL: ptr("https://example.com/path/to/image/phuket-sling"),
			Flavor:   ptr("Sweet"),
			Options: []menu.ItemOption{
				{
					Name:       "Cocktail",
					Category:   "Cocktail",
					ImageURL:   ptr("https://example.com/path/to/image/phuket-sling"),
					Materials:  []string{"Peach liqueur", "Blue curacao", "Grapefruit juice", "Tonic water"},
					OutOfStock: false,
					Price:      700,
				},
			},
		}
		assert.NoError(t, err)
		assert.Equal(t, want, got)
		menuSaveLister.AssertCalled(t, "Save", want)
		stockSaveLister.AssertNotCalled(t, "Save")
	})

	t.Run("will save new material", func(t *testing.T) {
		t.Parallel()

		recipeSaveLister := new(MockRecipeSaveListRemover)
		recipeGroups := func() iter.Seq2[recipe.RecipeGroup, error] {
			return func(yield func(recipe.RecipeGroup, error) bool) {
				for rg := range recipetest.ExampleRecipeGroupsIter {
					if !yield(rg, nil) {
						return
					}
				}
				yield(recipe.RecipeGroup{
					Name:     "New Recipe",
					ImageURL: ptr("https://example.com/path/to/image/new"),
					Recipes: []recipe.Recipe{
						{
							Name:  "New Recipe",
							Type:  "build",
							Glass: "collins",
							Steps: []recipe.Step{
								{
									Material: ptr("New Material"),
									Amount:   ptr("30ml"),
								},
							},
						},
					},
				}, nil)
			}
		}()
		recipeSaveLister.On("All").Return(recipeGroups, nil)

		menuSaveLister := new(MockMenu)
		menuSaveLister.On("Save", mock.Anything).Return(nil)

		stockSaveLister := new(MockStockSaveLister)
		stockSaveLister.On("All").Return(stocktest.ExampleMaterialsIter, nil)
		stockSaveLister.On("Save", mock.Anything, mock.Anything).Return(nil)

		c := &copilot{recipe: recipeSaveLister, menu: menuSaveLister, stock: stockSaveLister}
		got, err := c.SaveAsMenuItem("New Recipe", SaveAsMenuItemArg{
			Flavor: ptr("Fruity"),
			Options: map[string]MenuFromRecipeGroupArg{
				"New Recipe": {
					ImageURL: ptr("https://example.com/path/to/image/new-recipe"),
					Price:    700,
				},
			},
		})
		want := menu.Item{
			Name:     "New Recipe",
			ImageURL: ptr("https://example.com/path/to/image/new"),
			Flavor:   ptr("Fruity"),
			Options: []menu.ItemOption{
				{
					Name:       "New Recipe",
					ImageURL:   ptr("https://example.com/path/to/image/new-recipe"),
					Materials:  []string{"New Material"},
					OutOfStock: false,
					Price:      700,
				},
			},
		}
		assert.NoError(t, err)
		assert.Equal(t, want, got)
		menuSaveLister.AssertCalled(t, "Save", want)
		stockSaveLister.AssertCalled(t, "Save", []string{"New Material"}, mock.Anything)
	})

	t.Run("will call menu.Remove with recipeGroupName", func(t *testing.T) {
		t.Parallel()

		menuSaveLister := new(MockMenu)
		menuSaveLister.On("Remove", mock.Anything).Return(nil)

		c := &copilot{menu: menuSaveLister}
		_, err := c.SaveAsMenuItem("Phuket Sling", SaveAsMenuItemArg{Remove: true})
		assert.NoError(t, err)
		menuSaveLister.AssertCalled(t, "Remove", "Phuket Sling")
	})
}
