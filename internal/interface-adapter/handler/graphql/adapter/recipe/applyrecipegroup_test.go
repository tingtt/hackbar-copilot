package recipeadapter

import (
	"hackbar-copilot/internal/domain/recipe"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ApplyRecipeGroupTest struct {
	name  string
	base  recipe.RecipeGroup
	input model.InputRecipeGroup
	want  recipe.RecipeGroup
}

var applyRecipeGroupTests = []ApplyRecipeGroupTest{
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
			Recipes: []recipe.Recipe{{Name: "exists recipe", Category: "category 1"}},
		},
		input: model.InputRecipeGroup{
			Recipes: []*model.InputRecipe{{Name: "new recipe", Category: "category 1"}},
		},
		want: recipe.RecipeGroup{
			Name:    "ExistsRecipeGroup",
			Recipes: []recipe.Recipe{{Name: "exists recipe", Category: "category 1"}, {Name: "new recipe", Category: "category 1"}},
		},
	},
	{
		name: "will apply changed recipe",
		base: recipe.RecipeGroup{
			Name:    "ExistsRecipeGroup",
			Recipes: []recipe.Recipe{{Name: "exists recipe", Category: "category 1"}},
		},
		input: model.InputRecipeGroup{
			Recipes: []*model.InputRecipe{{
				Name:       "exists recipe",
				Category:   "category 1",
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
					Name:     "exists recipe",
					Category: "category 1",
					Type:     "recipe type",
					Glass:    "glass type",
					Steps: []recipe.Step{{
						Material:    ptr("whisky"),
						Amount:      ptr("30ml"),
						Description: ptr("description"),
					}},
				},
			},
		},
	},
	{
		name: "will apply changed recipe (replace)",
		base: recipe.RecipeGroup{
			Name:    "ExistsRecipeGroup",
			Recipes: []recipe.Recipe{{Name: "exists recipe", Category: "category 1"}},
		},
		input: model.InputRecipeGroup{
			Replace: ptr(true),
			Recipes: []*model.InputRecipe{{
				Name:     "new recipe",
				Category: "category 2",
			}},
		},
		want: recipe.RecipeGroup{
			Name: "ExistsRecipeGroup",
			Recipes: []recipe.Recipe{
				{
					Name:     "new recipe",
					Category: "category 2",
				},
			},
		},
	},
	{
		name: "will apply changed recipe (remove)",
		base: recipe.RecipeGroup{
			Name:    "ExistsRecipeGroup",
			Recipes: []recipe.Recipe{{Name: "exists recipe", Category: "category 1"}},
		},
		input: model.InputRecipeGroup{
			Recipes: []*model.InputRecipe{{
				Name:     "exists recipe",
				Category: "category 1",
				Remove:   ptr(true),
			}},
		},
		want: recipe.RecipeGroup{
			Name:    "ExistsRecipeGroup",
			Recipes: []recipe.Recipe{},
		},
	},
	{
		name: "will apply changed recipe (remove)",
		base: recipe.RecipeGroup{
			Name:    "ExistsRecipeGroup",
			Recipes: []recipe.Recipe{{Name: "exists recipe", Category: "category 1"}},
		},
		input: model.InputRecipeGroup{
			Replace: ptr(true),
			Recipes: []*model.InputRecipe{{
				Name:     "exists recipe",
				Category: "category 1",
				Remove:   ptr(true),
			}},
		},
		want: recipe.RecipeGroup{
			Name:    "ExistsRecipeGroup",
			Recipes: []recipe.Recipe{},
		},
	},
}

func Test_recipeAdapterIn_ApplyRecipeGroup(t *testing.T) {
	t.Parallel()

	for _, tt := range applyRecipeGroupTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := &inputAdapter{}
			got := s.ApplyRecipeGroup(tt.base, tt.input)

			assert.ElementsMatch(t, got.Recipes, tt.want.Recipes)
			got.Recipes = nil
			tt.want.Recipes = nil
			assert.Equal(t, tt.want, got)
		})
	}
}
