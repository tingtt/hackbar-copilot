package recipeadapter

import (
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	"hackbar-copilot/internal/usecase/copilot"
	"testing"

	"github.com/stretchr/testify/assert"
)

type applyAsMenuTest struct {
	name string
	arg  model.InputRecipeGroup
	want *copilot.SaveAsMenuItemArg
}

var applyAsMenuTests = []applyAsMenuTest{
	{
		name: "apply flavor",
		arg: model.InputRecipeGroup{
			AsMenu: &model.InputAsMenuArgs{
				Flavor: ptr("Sweet"),
			},
		},
		want: &copilot.SaveAsMenuItemArg{
			Flavor:  ptr("Sweet"),
			Options: map[string]copilot.MenuFromRecipeGroupArg{},
		},
	},
	{
		name: "apply items",
		arg: model.InputRecipeGroup{
			Recipes: []*model.InputRecipe{
				{
					Name: "Single",
					AsMenu: &model.InputAsMenuItemArgs{
						ImageURL: ptr("https://example.com/path/to/image/whisky-single"),
						Price:    1000,
					},
				},
				{
					Name: "Double",
					AsMenu: &model.InputAsMenuItemArgs{
						ImageURL: ptr("https://example.com/path/to/image/whisky-double"),
						Price:    2000,
					},
				},
			},
		},
		want: &copilot.SaveAsMenuItemArg{
			Flavor: nil,
			Options: map[string]copilot.MenuFromRecipeGroupArg{
				"Single": {
					ImageURL: ptr("https://example.com/path/to/image/whisky-single"),
					Price:    1000,
				},
				"Double": {
					ImageURL: ptr("https://example.com/path/to/image/whisky-double"),
					Price:    2000,
				},
			},
		},
	},
	{
		name: "apply remove",
		arg: model.InputRecipeGroup{
			AsMenu: &model.InputAsMenuArgs{
				Remove: ptr(true),
			},
		},
		want: &copilot.SaveAsMenuItemArg{Remove: true},
	},
}

func Test_recipeAdapterIn_ApplyAsMenu(t *testing.T) {
	t.Parallel()

	for _, tt := range applyAsMenuTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := &inputAdapter{}
			got, err := s.ApplyAsMenu(tt.arg)

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
