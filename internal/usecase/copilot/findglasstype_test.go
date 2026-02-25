package copilot

import (
	"hackbar-copilot/internal/domain/recipe"
	"hackbar-copilot/internal/domain/recipe/recipetest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_copilot_FindGlassType(t *testing.T) {
	t.Parallel()

	t.Run("will return mapped glass types", func(t *testing.T) {
		t.Parallel()

		glassTypes := recipetest.ExampleGlassTypesIter
		want := map[string]recipe.GlassType{
			"collins": {
				Name:        "collins",
				ImageURL:    new("https://example.com/path/to/image/collins"),
				Description: new("collins glass description"),
			},
			"cocktail": {
				Name:        "cocktail",
				ImageURL:    new("https://example.com/path/to/image/cocktail"),
				Description: new("cocktail glass description"),
			},
			"shot": {
				Name:        "shot",
				ImageURL:    new("https://example.com/path/to/image/shot"),
				Description: new("shot glass description"),
			},
			"rock": {
				Name:        "rock",
				ImageURL:    new("https://example.com/path/to/image/rock"),
				Description: new("rock glass description"),
			},
			"beer": {
				Name:        "beer",
				ImageURL:    new("https://example.com/path/to/image/beer"),
				Description: new("beer glass description"),
			},
		}

		recipeMock := new(MockRecipe)
		recipeMock.On("AllGlassTypes").Return(glassTypes, nil)
		gateway := MockGateway{recipe: recipeMock}

		c := &copilot{&gateway}

		got, err := c.FindGlassType()
		assert.NoError(t, err)
		assert.Equal(t, want, got)
	})
}
