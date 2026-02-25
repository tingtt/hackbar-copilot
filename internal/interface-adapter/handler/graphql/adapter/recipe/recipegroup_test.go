package recipeadapter

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
			Description: new("build description"),
		},
	}
	recipeTypesModel = map[string]*model.RecipeType{
		"build": {
			Name:        "build",
			Description: new("build description"),
		},
	}
	glassTypes = map[string]recipe.GlassType{
		"collins": {
			Name:        "collins",
			ImageURL:    new("https://example.com/path/to/image"),
			Description: new("collins glass description"),
		},
	}
	glassTypesModel = map[string]*model.GlassType{
		"collins": {
			Name:        "collins",
			ImageURL:    new("https://example.com/path/to/image"),
			Description: new("collins glass description"),
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
			ImageURL: new("https://example.com/path/to/image"),
			Recipes: []recipe.Recipe{
				{
					Name: "Recipe name",
					Steps: []recipe.Step{
						{
							Material: new("Peach liqueur"),
							Amount:   new("30ml"),
						},
						{
							Material: new("Blue curacao"),
							Amount:   new("15ml"),
						},
						{
							Material: new("Grapefruit juice"),
							Amount:   new("30ml"),
						},
						{
							Description: new("Stir"),
						},
						{
							Material: new("Tonic water"),
							Amount:   new("Full up"),
						},
					},
				},
			},
		},
		out: &model.RecipeGroup{
			Name:     "RecipeGroup name",
			ImageURL: new("https://example.com/path/to/image"),
			Recipes: []*model.Recipe{
				{
					Name: "Recipe name",
					Steps: []*model.Step{
						{
							Material: new("Peach liqueur"),
							Amount:   new("30ml"),
						},
						{
							Material: new("Blue curacao"),
							Amount:   new("15ml"),
						},
						{
							Material: new("Grapefruit juice"),
							Amount:   new("30ml"),
						},
						{
							Description: new("Stir"),
						},
						{
							Material: new("Tonic water"),
							Amount:   new("Full up"),
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
			ImageURL: new("https://example.com/path/to/image"),
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
			ImageURL: new("https://example.com/path/to/image"),
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
			a := &outputAdapter{}
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
				a := &outputAdapter{}
				if got := a.recipe(recipeTypes, glassTypes)(r); !reflect.DeepEqual(*got, *tt.out.Recipes[i]) {
					t.Errorf("adapterOut.recipe() = %v, want %v", got, tt.out.Recipes[i])
				}
			}
		})
	}
}
