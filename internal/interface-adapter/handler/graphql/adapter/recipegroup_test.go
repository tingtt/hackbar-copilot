package adapter

import (
	"hackbar-copilot/internal/domain/recipe"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	"reflect"
	"testing"
)

var (
	recipeTypes = map[string]recipe.RecipeType{
		"build": {
			Name:        "build",
			Description: ptr("build description"),
		},
	}
	recipeTypesModel = map[string]*model.RecipeType{
		"build": {
			Name:        "build",
			Description: ptr("build description"),
		},
	}
	glassTypes = map[string]recipe.GlassType{
		"collins": {
			Name:        "collins",
			ImageURL:    ptr("https://example.com/path/to/image"),
			Description: ptr("collins glass description"),
		},
	}
	glassTypesModel = map[string]*model.GlassType{
		"collins": {
			Name:        "collins",
			ImageURL:    ptr("https://example.com/path/to/image"),
			Description: ptr("collins glass description"),
		},
	}
)

type recipeGroupTest struct {
	name string
	in   recipe.RecipeGroup
	out  *model.RecipeGroup
}

var recipeGroupTests = []recipeGroupTest{
	{
		name: "will adapt RecipeGroup",
		in: recipe.RecipeGroup{
			Name:     "RecipeGroup name",
			ImageURL: ptr("https://example.com/path/to/image"),
			Recipes: []recipe.Recipe{
				{
					Name: "Recipe name",
					Steps: []recipe.Step{
						{
							Material: ptr("Peach liqueur"),
							Amount:   ptr("30ml"),
						},
						{
							Material: ptr("Blue curacao"),
							Amount:   ptr("15ml"),
						},
						{
							Material: ptr("Grapefruit juice"),
							Amount:   ptr("30ml"),
						},
						{
							Description: ptr("Stir"),
						},
						{
							Material: ptr("Tonic water"),
							Amount:   ptr("Full up"),
						},
					},
				},
			},
		},
		out: &model.RecipeGroup{
			Name:     "RecipeGroup name",
			ImageURL: ptr("https://example.com/path/to/image"),
			Recipes: []*model.Recipe{
				{
					Name: "Recipe name",
					Steps: []*model.Step{
						{
							Material: ptr("Peach liqueur"),
							Amount:   ptr("30ml"),
						},
						{
							Material: ptr("Blue curacao"),
							Amount:   ptr("15ml"),
						},
						{
							Material: ptr("Grapefruit juice"),
							Amount:   ptr("30ml"),
						},
						{
							Description: ptr("Stir"),
						},
						{
							Material: ptr("Tonic water"),
							Amount:   ptr("Full up"),
						},
					},
				},
			},
		},
	},
	{
		name: "will adapt RecipeGroup with matched recipeType and GlassType",
		in: recipe.RecipeGroup{
			Name:     "Phuket Sling",
			ImageURL: ptr("https://example.com/path/to/image"),
			Recipes: []recipe.Recipe{
				{
					Name:  "Cocktail",
					Type:  "build",
					Glass: "collins",
					Steps: []recipe.Step{},
				},
			},
		},
		out: &model.RecipeGroup{
			Name:     "Phuket Sling",
			ImageURL: ptr("https://example.com/path/to/image"),
			Recipes: []*model.Recipe{
				{
					Name:  "Cocktail",
					Type:  recipeTypesModel["build"],
					Glass: glassTypesModel["collins"],
					Steps: []*model.Step{},
				},
			},
		},
	},
	{
		name: "will adapt RecipeGroup without matched recipeType and GlassType",
		in: recipe.RecipeGroup{
			Name: "Phuket Sling",
			Recipes: []recipe.Recipe{
				{
					Name:  "Cocktail",
					Type:  "not exists recipe type",
					Glass: "not exists glass type",
					Steps: []recipe.Step{},
				},
			},
		},
		out: &model.RecipeGroup{
			Name: "Phuket Sling",
			Recipes: []*model.Recipe{
				{
					Name:  "Cocktail",
					Type:  nil,
					Glass: nil,
					Steps: []*model.Step{},
				},
			},
		},
	},
}

func Test_adapterOut_RecipeGroup(t *testing.T) {
	t.Parallel()

	for _, tt := range recipeGroupTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := &adapterOut{}
			if got := a.RecipeGroup(recipeTypes, glassTypes)(tt.in); !reflect.DeepEqual(*got, *tt.out) {
				t.Errorf("adapterOut.RecipeGroup() = %v, want %v", got, tt.out)
			}
		})
	}
}

func Test_adapterOut_recipe(t *testing.T) {
	t.Parallel()

	for _, tt := range recipeGroupTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			for i, r := range tt.in.Recipes {
				a := &adapterOut{}
				if got := a.recipe(recipeTypes, glassTypes)(r); !reflect.DeepEqual(*got, *tt.out.Recipes[i]) {
					t.Errorf("adapterOut.recipe() = %v, want %v", got, tt.out.Recipes[i])
				}
			}
		})
	}
}
