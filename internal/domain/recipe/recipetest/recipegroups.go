package recipetest

import (
	"hackbar-copilot/internal/domain/recipe"
)

var ExampleRecipeGroups = []recipe.RecipeGroup{
	{
		Name:     "Phuket Sling",
		ImageURL: ptr("https://example.com/path/to/image/phuket-sling"),
		Recipes: []recipe.Recipe{
			{
				Name:  "Cocktail",
				Type:  "build",
				Glass: "collins",
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
	{
		Name:     "Passoamoni",
		ImageURL: ptr("https://example.com/path/to/image/passoamoni"),
		Recipes: []recipe.Recipe{
			{
				Name:  "Cocktail",
				Type:  "build",
				Glass: "collins",
				Steps: []recipe.Step{
					{
						Material: ptr("Passoa"),
						Amount:   ptr("45ml"),
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
	{
		Name:     "Blue Devil",
		ImageURL: ptr("https://example.com/path/to/image/passoamoni"),
		Recipes: []recipe.Recipe{
			{
				Name:  "Cocktail",
				Type:  "shake",
				Glass: "cocktail",
				Steps: []recipe.Step{
					{
						Description: ptr("Chill shaker and glass."),
					},
					{
						Description: ptr("Put ingredients in a shaker."),
					},
					{
						Material: ptr("Gin"),
						Amount:   ptr("30ml"),
					},
					{
						Material: ptr("Blue curacao"),
						Amount:   ptr("15ml"),
					},
					{
						Material: ptr("Lemon juice"),
						Amount:   ptr("15ml"),
					},
					{
						Description: ptr("Put ice in a shaker."),
					},
					{
						Description: ptr("Shake."),
					},
					{
						Description: ptr("Pour into a glass."),
					},
				},
			},
		},
	},
}

var ExampleRecipeGroupsIter = iterWithNilError(ExampleRecipeGroups)
var ExampleRecipeGroupsMap = func() map[string]recipe.RecipeGroup {
	m := make(map[string]recipe.RecipeGroup)
	for _, v := range ExampleRecipeGroups {
		m[v.Name] = v
	}
	return m
}()
