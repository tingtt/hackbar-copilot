package recipetest

import (
	"hackbar-copilot/internal/domain/recipe"
)

var ExampleGlassTypes = []recipe.GlassType{
	{
		Name:        "collins",
		ImageURL:    new("https://example.com/path/to/image/collins"),
		Description: new("collins glass description"),
	},
	{
		Name:        "cocktail",
		ImageURL:    new("https://example.com/path/to/image/cocktail"),
		Description: new("cocktail glass description"),
	},
	{
		Name:        "shot",
		ImageURL:    new("https://example.com/path/to/image/shot"),
		Description: new("shot glass description"),
	},
	{
		Name:        "rock",
		ImageURL:    new("https://example.com/path/to/image/rock"),
		Description: new("rock glass description"),
	},
	{
		Name:        "beer",
		ImageURL:    new("https://example.com/path/to/image/beer"),
		Description: new("beer glass description"),
	},
}

var ExampleGlassTypesIter = IterWithNilErrorGlassTypes(ExampleGlassTypes)
var ExampleGlassTypesMap = func() map[string]recipe.GlassType {
	m := make(map[string]recipe.GlassType)
	for _, v := range ExampleGlassTypes {
		m[v.Name] = v
	}
	return m
}()
