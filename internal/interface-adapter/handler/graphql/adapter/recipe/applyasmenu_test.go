package recipeadapter

import (
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	"hackbar-copilot/internal/usecase/copilot"
	"testing"

	"github.com/stretchr/testify/assert"
)

type applyAsMenuTest struct {
	name    string
	arg     model.InputRecipeGroup
	want    *copilot.SaveAsMenuItemArg
	wantErr bool
}

var applyAsMenuTests = []applyAsMenuTest{
	{
		name: "nil",
		arg: model.InputRecipeGroup{
			AsMenu: nil,
		},
		want: nil,
	},
	{
		name: "apply flavor",
		arg: model.InputRecipeGroup{
			AsMenu: &model.InputAsMenuItemArgs{
				Flavor: new("Sweet"),
			},
		},
		want: &copilot.SaveAsMenuItemArg{
			Flavor:  new("Sweet"),
			Options: map[string]copilot.MenuFromRecipeGroupArg{},
		},
	},
	{
		name: "apply items",
		arg: model.InputRecipeGroup{
			Recipes: []*model.InputRecipe{
				{
					Name: "Single",
					AsMenu: &model.InputAsMenuItemOptionArgs{
						ImageURL: new("https://example.com/path/to/image/whisky-single"),
						Price:    1000,
					},
				},
				{
					Name: "Double",
					AsMenu: &model.InputAsMenuItemOptionArgs{
						ImageURL: new("https://example.com/path/to/image/whisky-double"),
						Price:    2000,
					},
				},
			},
			AsMenu: &model.InputAsMenuItemArgs{Flavor: nil},
		},
		want: &copilot.SaveAsMenuItemArg{
			Flavor: nil,
			Options: map[string]copilot.MenuFromRecipeGroupArg{
				"Single": {
					ImageURL: new("https://example.com/path/to/image/whisky-single"),
					Price:    1000,
				},
				"Double": {
					ImageURL: new("https://example.com/path/to/image/whisky-double"),
					Price:    2000,
				},
			},
		},
	},
	{
		name: "apply remove",
		arg: model.InputRecipeGroup{
			AsMenu: &model.InputAsMenuItemArgs{
				Remove: new(true),
			},
		},
		want: &copilot.SaveAsMenuItemArg{Remove: true},
	},
	{
		name: "apply items/input.AsMenu not specified",
		arg: model.InputRecipeGroup{
			Recipes: []*model.InputRecipe{
				{
					Name: "Single",
					AsMenu: &model.InputAsMenuItemOptionArgs{
						ImageURL: new("https://example.com/path/to/image/whisky-single"),
						Price:    1000,
					},
				},
				{
					Name: "Double",
					AsMenu: &model.InputAsMenuItemOptionArgs{
						ImageURL: new("https://example.com/path/to/image/whisky-double"),
						Price:    2000,
					},
				},
			},
			AsMenu: nil,
		},
		wantErr: true,
	},
}

func Test_recipeAdapterIn_ApplyAsMenu(t *testing.T) {
	t.Parallel()

	for _, tt := range applyAsMenuTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := &inputAdapter{}
			got, err := s.ApplyAsMenu(tt.arg)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
