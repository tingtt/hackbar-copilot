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
			"shake": {Name: "shake", Description: ptr("shake description")},
			"build": {Name: "build", Description: ptr("build description")},
			"stir":  {Name: "stir", Description: ptr("stir description")},
			"blend": {Name: "blend", Description: ptr("blend description")},
		}

		recipeSaveLister := new(MockRecipeSaveListRemover)
		recipeSaveLister.On("AllRecipeTypes").Return(recipeTypes, nil)

		c := &copilot{recipe: recipeSaveLister}

		got, err := c.FindRecipeType()
		assert.NoError(t, err)
		assert.Equal(t, want, got)
	})
}
