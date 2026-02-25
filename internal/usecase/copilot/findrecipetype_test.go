package copilot

import (
	"hackbar-copilot/internal/domain/recipe"
	"hackbar-copilot/internal/domain/recipe/recipetest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_copilot_FindRecipeType(t *testing.T) {
	t.Parallel()

	t.Run("will return mapped recipe types", func(t *testing.T) {
		t.Parallel()

		recipeTypes := recipetest.ExampleRecipeTypesIter
		want := map[string]recipe.RecipeType{
			"shake": {Name: "shake", Description: new("shake description")},
			"build": {Name: "build", Description: new("build description")},
			"stir":  {Name: "stir", Description: new("stir description")},
			"blend": {Name: "blend", Description: new("blend description")},
		}

		recipeMock := new(MockRecipe)
		recipeMock.On("AllRecipeTypes").Return(recipeTypes, nil)
		gateway := MockGateway{recipe: recipeMock}

		c := &copilot{&gateway}

		got, err := c.FindRecipeType()
		assert.NoError(t, err)
		assert.Equal(t, want, got)
	})
}
