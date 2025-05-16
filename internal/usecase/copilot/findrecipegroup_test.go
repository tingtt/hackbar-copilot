package copilot

import (
	"hackbar-copilot/internal/domain/recipe/recipetest"
	usecaseutils "hackbar-copilot/internal/usecase/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_copilot_FindRecipeGroup(t *testing.T) {
	t.Parallel()

	t.Run("will return recipe group", func(t *testing.T) {
		t.Parallel()

		recipeGroup := recipetest.ExampleRecipeGroupsIter
		want := recipetest.ExampleRecipeGroups[0]

		recipeMock := new(MockRecipe)
		recipeMock.On("All").Return(recipeGroup, nil)
		gateway := MockGateway{recipe: recipeMock}

		c := &copilot{&gateway}

		got, err := c.FindRecipeGroup(want.Name)
		assert.NoError(t, err)
		assert.Equal(t, want, got)
	})

	t.Run("will return not found error", func(t *testing.T) {
		t.Parallel()

		recipeGroup := recipetest.ExampleRecipeGroupsIter

		recipeMock := new(MockRecipe)
		recipeMock.On("All").Return(recipeGroup, nil)
		gateway := MockGateway{recipe: recipeMock}

		c := &copilot{&gateway}

		_, err := c.FindRecipeGroup("notfound")
		assert.ErrorIs(t, err, usecaseutils.ErrNotFound)
	})
}
