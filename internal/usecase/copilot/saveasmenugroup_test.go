package copilot

import (
	"hackbar-copilot/internal/domain/menu"
	"hackbar-copilot/internal/domain/recipe/recipetest"
	usecaseutils "hackbar-copilot/internal/usecase/utils"
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

			menuSaveLister := new(MockMenuSaveLister)
			menuSaveLister.On("Save", mock.Anything).Return(nil)

			c := &copilot{recipe: recipeSaveLister, menu: menuSaveLister}
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

		menuSaveLister := new(MockMenuSaveLister)
		menuSaveLister.On("Save", mock.Anything).Return(nil)

		c := &copilot{recipe: recipeSaveLister, menu: menuSaveLister}
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
	})
}
