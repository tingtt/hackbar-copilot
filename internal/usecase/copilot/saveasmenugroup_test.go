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

func Test_copilot_SaveAsMenuGroup(t *testing.T) {
	t.Parallel()

	t.Run("will call copilot.FindRecipeGroup with recipeGroupName", func(t *testing.T) {
		t.Parallel()

		t.Run("found", func(t *testing.T) {
			t.Parallel()

			recipeSaveLister := new(MockRecipeSaveLister)
			recipeSaveLister.On("All").Return(recipetest.ExampleRecipeGroupsIter, nil)

			menuSaveLister := new(MockMenu)
			menuSaveLister.On("Save", mock.Anything).Return(nil)

			stockSaveLister := new(MockStockSaveLister)
			stockSaveLister.On("All").Return(stocktest.ExampleMaterialsIter, nil)
			stockSaveLister.On("Save", mock.Anything, mock.Anything).Return(nil)

			c := &copilot{recipe: recipeSaveLister, menu: menuSaveLister, stock: stockSaveLister}
			got, err := c.SaveAsMenuGroup("Phuket Sling", SaveAsMenuGroupArg{})
			assert.NoError(t, err)
			assert.Equal(t, "Phuket Sling", got.Name)
		})

		t.Run("not found", func(t *testing.T) {
			t.Parallel()

			recipeSaveLister := new(MockRecipeSaveLister)
			recipeSaveLister.On("All").Return(recipetest.ExampleRecipeGroupsIter, nil)

			c := &copilot{recipe: recipeSaveLister}
			_, err := c.SaveAsMenuGroup("-", SaveAsMenuGroupArg{})
			assert.ErrorIs(t, err, usecaseutils.ErrNotFound)
		})
	})

	t.Run("will return menu group converted from recipe group", func(t *testing.T) {
		t.Parallel()

		recipeSaveLister := new(MockRecipeSaveLister)
		recipeSaveLister.On("All").Return(recipetest.ExampleRecipeGroupsIter, nil)

		menuSaveLister := new(MockMenu)
		menuSaveLister.On("Save", mock.Anything).Return(nil)

		stockSaveLister := new(MockStockSaveLister)
		stockSaveLister.On("All").Return(stocktest.ExampleMaterialsIter, nil)

		c := &copilot{recipe: recipeSaveLister, menu: menuSaveLister, stock: stockSaveLister}
		got, err := c.SaveAsMenuGroup("Phuket Sling", SaveAsMenuGroupArg{
			Flavor: ptr("Sweet"),
			Items: map[string]MenuFromRecipeGroupArg{
				"Cocktail": {
					ImageURL: ptr("https://example.com/path/to/image/phuket-sling"),
					Price:    700,
				},
			},
		})
		want := menu.Group{
			Name:     "Phuket Sling",
			ImageURL: ptr("https://example.com/path/to/image/phuket-sling"),
			Flavor:   ptr("Sweet"),
			Items: []menu.Item{
				{
					Name:       "Cocktail",
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

		recipeSaveLister := new(MockRecipeSaveLister)
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
		got, err := c.SaveAsMenuGroup("New Recipe", SaveAsMenuGroupArg{
			Flavor: ptr("Fruity"),
			Items: map[string]MenuFromRecipeGroupArg{
				"New Recipe": {
					ImageURL: ptr("https://example.com/path/to/image/new-recipe"),
					Price:    700,
				},
			},
		})
		want := menu.Group{
			Name:     "New Recipe",
			ImageURL: ptr("https://example.com/path/to/image/new"),
			Flavor:   ptr("Fruity"),
			Items: []menu.Item{
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
}
