package recipetest

import (
	"hackbar-copilot/internal/domain/recipe"
)

var ExampleGlassTypes = []recipe.GlassType{
	{
		Name:        "collins",
		ImageURL:    ptr("https://example.com/path/to/image/collins"),
		Description: ptr("collins glass description"),
	},
	{
		Name:        "cocktail",
		ImageURL:    ptr("https://example.com/path/to/image/cocktail"),
		Description: ptr("cocktail glass description"),
	},
	{
		Name:        "shot",
		ImageURL:    ptr("https://example.com/path/to/image/shot"),
		Description: ptr("shot glass description"),
	},
	{
		Name:        "rock",
		ImageURL:    ptr("https://example.com/path/to/image/rock"),
		Description: ptr("rock glass description"),
	},
	{
		Name:        "beer",
		ImageURL:    ptr("https://example.com/path/to/image/beer"),
		Description: ptr("beer glass description"),
	},
}

var ExampleGlassTypesIter = iterWithNilError(ExampleGlassTypes)
var ExampleGlassTypesMap = func() map[string]recipe.GlassType {
	m := make(map[string]recipe.GlassType)
	for _, v := range ExampleGlassTypes {
		m[v.Name] = v
	}
	return m
}()
