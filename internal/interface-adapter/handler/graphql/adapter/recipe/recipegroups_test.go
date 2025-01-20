package recipeadapter

import (
	"hackbar-copilot/internal/domain/recipe"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	"hackbar-copilot/internal/utils/sliceutil"
	"reflect"
	"testing"
)

func Test_adapterOut_RecipeGroups(t *testing.T) {
	t.Parallel()

	t.Run("will return adapted recipe groups", func(t *testing.T) {
		t.Parallel()

		in := sliceutil.Map(recipeGroupTests, func(recipeGroupTest recipeGroupTest) recipe.RecipeGroup {
			return recipeGroupTest.in
		})
		want := sliceutil.Map(recipeGroupTests, func(recipeGroupTest recipeGroupTest) *model.RecipeGroup {
			return recipeGroupTest.out
		})

		adapter := &outputAdapter{}
		got := adapter.RecipeGroups(in, recipeTypes, glassTypes)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("adapterOut.RecipeGroups() = %v, want %v", got, want)
		}
	})
}
