package recipetest

import (
	"hackbar-copilot/internal/domain/recipe"
)

var ExampleRecipeTypes = []recipe.RecipeType{
	{
		Name:        "shake",
		Description: new("shake description"),
	},
	{
		Name:        "build",
		Description: new("build description"),
	},
	{
		Name:        "stir",
		Description: new("stir description"),
	},
	{
		Name:        "blend",
		Description: new("blend description"),
	},
}

var ExampleRecipeTypesIter = IterWithNilErrorRecipeTypes(ExampleRecipeTypes)
var ExampleRecipeTypesMap = func() map[string]recipe.RecipeType {
	m := make(map[string]recipe.RecipeType)
	for _, v := range ExampleRecipeTypes {
		m[v.Name] = v
	}
	return m
}()
