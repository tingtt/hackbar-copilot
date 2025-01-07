package adapter

import (
	"hackbar-copilot/internal/domain/recipe"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	"reflect"
	"testing"
)

type ApplyRecipeGroupTest struct {
	name  string
	base  recipe.RecipeGroup
	input model.InputRecipeGroup
	want  recipe.RecipeGroup
}

var testApplyRecipeGroup = []ApplyRecipeGroupTest{
	{
		name:  "will apply name",
		base:  recipe.RecipeGroup{},
		input: model.InputRecipeGroup{Name: "NewRecipeGroup"},
		want:  recipe.RecipeGroup{Name: "NewRecipeGroup"},
	},
	{
		name: "will apply image URL",
		base: recipe.RecipeGroup{Name: "ExistsRecipeGroup"},
		input: model.InputRecipeGroup{
			Name:     "ExistsRecipeGroup",
			ImageURL: ptr("https://example.com/path/to/image"),
		},
		want: recipe.RecipeGroup{
			Name:     "ExistsRecipeGroup",
			ImageURL: ptr("https://example.com/path/to/image"),
		},
	},
	{
		name: "will apply new recipe",
		base: recipe.RecipeGroup{
			Name:    "ExistsRecipeGroup",
			Recipes: []recipe.Recipe{{Name: "exists recipe"}},
		},
		input: model.InputRecipeGroup{
			Recipes: []*model.InputRecipe{{Name: "new recipe"}},
		},
		want: recipe.RecipeGroup{
			Name:    "ExistsRecipeGroup",
			Recipes: []recipe.Recipe{{Name: "exists recipe"}, {Name: "new recipe"}},
		},
	},
	{
		name: "will apply changed recipe",
		base: recipe.RecipeGroup{
			Name:    "ExistsRecipeGroup",
			Recipes: []recipe.Recipe{{Name: "exists recipe"}},
		},
		input: model.InputRecipeGroup{
			Recipes: []*model.InputRecipe{{
				Name:       "exists recipe",
				RecipeType: &model.InputRecipeType{Name: "recipe type"},
				GlassType:  &model.InputGlassType{Name: "glass type"},
				Steps: []*model.InputStep{{
					Material:    ptr("whisky"),
					Amount:      ptr("30ml"),
					Description: ptr("description"),
				}},
			}},
		},
		want: recipe.RecipeGroup{
			Name: "ExistsRecipeGroup",
			Recipes: []recipe.Recipe{
				{
					Name:  "exists recipe",
					Type:  "recipe type",
					Glass: "glass type",
					Steps: []recipe.Step{{
						Material:    ptr("whisky"),
						Amount:      ptr("30ml"),
						Description: ptr("description"),
					}},
				},
			},
		},
	},
}

func Test_recipeAdapterIn_ApplyRecipeGroup(t *testing.T) {
	t.Parallel()

	for _, tt := range testApplyRecipeGroup {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := &recipeAdapterIn{}
			if got := s.ApplyRecipeGroup(tt.base, tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("recipeAdapterIn.ApplyRecipeGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
