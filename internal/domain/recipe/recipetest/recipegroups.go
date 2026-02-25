package recipetest

import (
	"hackbar-copilot/internal/domain/recipe"
)

var ExampleRecipeGroups = []recipe.RecipeGroup{
	{
		Name:     "Phuket Sling",
		ImageURL: new("https://example.com/path/to/image/phuket-sling"),
		Recipes: []recipe.Recipe{
			{
				Name:     "Cocktail",
				Category: "Cocktail",
				Type:     "build",
				Glass:    "collins",
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
	{
		Name:     "Passoamoni",
		ImageURL: new("https://example.com/path/to/image/passoamoni"),
		Recipes: []recipe.Recipe{
			{
				Name:     "Cocktail",
				Category: "Cocktail",
				Type:     "build",
				Glass:    "collins",
				Steps: []recipe.Step{
					{
						Material: new("Passoa"),
						Amount:   new("45ml"),
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
	{
		Name:     "Blue Devil",
		ImageURL: new("https://example.com/path/to/image/passoamoni"),
		Recipes: []recipe.Recipe{
			{
				Name:     "Cocktail",
				Category: "Cocktail",
				Type:     "shake",
				Glass:    "cocktail",
				Steps: []recipe.Step{
					{
						Description: new("Chill shaker and glass."),
					},
					{
						Description: new("Put ingredients in a shaker."),
					},
					{
						Material: new("Gin"),
						Amount:   new("30ml"),
					},
					{
						Material: new("Blue curacao"),
						Amount:   new("15ml"),
					},
					{
						Material: new("Lemon juice"),
						Amount:   new("15ml"),
					},
					{
						Description: new("Put ice in a shaker."),
					},
					{
						Description: new("Shake."),
					},
					{
						Description: new("Pour into a glass."),
					},
				},
			},
		},
	},
}

var ExampleRecipeGroupsIter = IterWithNilError(ExampleRecipeGroups)
var ExampleRecipeGroupsMap = func() map[string]recipe.RecipeGroup {
	m := make(map[string]recipe.RecipeGroup)
	for _, v := range ExampleRecipeGroups {
		m[v.Name] = v
	}
	return m
}()
