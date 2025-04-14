package copilot

import (
	"hackbar-copilot/internal/domain/recipe/recipetest"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_copilot_ListRecipes(t *testing.T) {
	t.Parallel()

	t.Run("will return sorted recipe groups", func(t *testing.T) {
		t.Parallel()

		recipeGroup := recipetest.ExampleRecipeGroupsIter
		want := recipetest.ExampleRecipeGroups
		sort.Slice(want, func(i, j int) bool {
			return want[i].Name < want[j].Name
		})

		recipeSaveLister := new(MockRecipeSaveListRemover)
		recipeSaveLister.On("All").Return(recipeGroup, nil)

		c := &copilot{recipe: recipeSaveLister}

		got, err := c.ListRecipes(SortRecipeGroupByName())
		assert.NoError(t, err)
		assert.Equal(t, want, got)
	})
}
