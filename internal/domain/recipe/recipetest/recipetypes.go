package recipetest

import (
	"hackbar-copilot/internal/domain/recipe"
)

var ExampleRecipeTypes = []recipe.RecipeType{
	{
		Name:        "shake",
		Description: ptr("shake description"),
	},
	{
		Name:        "build",
		Description: ptr("build description"),
	},
	{
		Name:        "stir",
		Description: ptr("stir description"),
	},
	{
		Name:        "blend",
		Description: ptr("blend description"),
	},
}

var ExampleRecipeTypesIter = iterWithNilError(ExampleRecipeTypes)
var ExampleRecipeTypesMap = func() map[string]recipe.RecipeType {
	m := make(map[string]recipe.RecipeType)
	for _, v := range ExampleRecipeTypes {
		m[v.Name] = v
	}
	return m
}()
