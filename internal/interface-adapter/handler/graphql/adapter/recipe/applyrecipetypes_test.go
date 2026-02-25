package recipeadapter

import (
	"hackbar-copilot/internal/domain/recipe"
	"hackbar-copilot/internal/domain/recipe/recipetest"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/adapter"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ApplyRecipeTypesTest struct {
	name    string
	arg     model.InputRecipeGroup
	want    []recipe.RecipeType
	wantErr error
}

var currentRecipeTypes = recipetest.ExampleRecipeTypesMap

var applyRecipeTypesTests = []ApplyRecipeTypesTest{
	{
		name: "append new recipe type",
		arg: model.InputRecipeGroup{
			Name: "recipe group",
			Recipes: []*model.InputRecipe{
				{
					Name: "recipe",
					RecipeType: &model.InputRecipeType{
						Name:        "new recipe type",
						Description: new("new recipe type description"),
						Save:        nil,
					},
				},
			},
		},
		want: []recipe.RecipeType{
			{
				Name:        "new recipe type",
				Description: new("new recipe type description"),
			},
		},
	},
	{
		name: "update with save flag",
		arg: model.InputRecipeGroup{
			Name: "recipe group",
			Recipes: []*model.InputRecipe{
				{
					Name: "recipe",
					RecipeType: &model.InputRecipeType{
						Name:        "collins",
						Description: new("exists recipe type description"),
						Save:        new(true),
					},
				},
			},
		},
		want: []recipe.RecipeType{
			{
				Name:        "collins",
				Description: new("exists recipe type description"),
			},
		},
	},
	{
		name: "update without save flag",
		arg: model.InputRecipeGroup{
			Name: "recipe group",
			Recipes: []*model.InputRecipe{
				{
					Name: "recipe",
					RecipeType: &model.InputRecipeType{
						Name:        "build",
						Description: new("exists recipe type description"),
						Save:        new(false),
					},
				},
			},
		},
		wantErr: adapter.ErrSaveDoesNotExist,
	},
	{
		name: "no updates",
		arg: model.InputRecipeGroup{
			Name: "recipe group",
			Recipes: []*model.InputRecipe{
				{
					Name:       "recipe",
					RecipeType: &model.InputRecipeType{Name: "build"},
				},
			},
		},
		want: []recipe.RecipeType{},
	},
}

func Test_recipeAdapterIn_ApplyRecipeTypes(t *testing.T) {
	t.Parallel()

	for _, tt := range applyRecipeTypesTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := &inputAdapter{}
			got, err := s.ApplyRecipeTypes(currentRecipeTypes, tt.arg)

			assert.Equal(t, tt.want, got)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}
