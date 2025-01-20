package recipeadapter

import (
	"hackbar-copilot/internal/domain/recipe"
	"hackbar-copilot/internal/domain/recipe/recipetest"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/adapter"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ApplyGlassTypesTest struct {
	name    string
	arg     model.InputRecipeGroup
	want    []recipe.GlassType
	wantErr error
}

var currentGlassTypes = recipetest.ExampleGlassTypesMap

var applyGlassTypesTests = []ApplyGlassTypesTest{
	{
		name: "append new glass type",
		arg: model.InputRecipeGroup{
			Name: "recipe group",
			Recipes: []*model.InputRecipe{
				{
					Name: "recipe",
					GlassType: &model.InputGlassType{
						Name:        "new glass type",
						ImageURL:    ptr("https://example.com/path/to/image/new-glass-type"),
						Description: ptr("new glass type description"),
						Save:        nil,
					},
				},
			},
		},
		want: []recipe.GlassType{
			{
				Name:        "new glass type",
				ImageURL:    ptr("https://example.com/path/to/image/new-glass-type"),
				Description: ptr("new glass type description"),
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
					GlassType: &model.InputGlassType{
						Name:        "collins",
						ImageURL:    ptr("https://example.com/path/to/image/exists-glass-type"),
						Description: ptr("exists glass type description"),
						Save:        ptr(true),
					},
				},
			},
		},
		want: []recipe.GlassType{
			{
				Name:        "collins",
				ImageURL:    ptr("https://example.com/path/to/image/exists-glass-type"),
				Description: ptr("exists glass type description"),
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
					GlassType: &model.InputGlassType{
						Name:        "collins",
						ImageURL:    ptr("https://example.com/path/to/image/exists-glass-type"),
						Description: ptr("exists glass type description"),
						Save:        ptr(false),
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
					Name:      "recipe",
					GlassType: &model.InputGlassType{Name: "collins"},
				},
			},
		},
		want: []recipe.GlassType{},
	},
}

func Test_recipeAdapterIn_ApplyGlassTypes(t *testing.T) {
	t.Parallel()

	for _, tt := range applyGlassTypesTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := &inputAdapter{}
			got, err := s.ApplyGlassTypes(currentGlassTypes, tt.arg)

			assert.Equal(t, tt.want, got)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}
