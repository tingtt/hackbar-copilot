package copilot

import (
	"hackbar-copilot/internal/domain/recipe"
	"hackbar-copilot/internal/domain/recipe/recipetest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ptr[T any](v T) *T {
	return &v
}

func Test_copilot_FindGlassType(t *testing.T) {
	t.Parallel()

	t.Run("will return mapped glass types", func(t *testing.T) {
		t.Parallel()

		glassTypes := recipetest.ExampleGlassTypesIter
		want := map[string]recipe.GlassType{
			"collins": {
				Name:        "collins",
				ImageURL:    ptr("https://example.com/path/to/image/collins"),
				Description: ptr("collins glass description"),
			},
			"cocktail": {
				Name:        "cocktail",
				ImageURL:    ptr("https://example.com/path/to/image/cocktail"),
				Description: ptr("cocktail glass description"),
			},
			"shot": {
				Name:        "shot",
				ImageURL:    ptr("https://example.com/path/to/image/shot"),
				Description: ptr("shot glass description"),
			},
			"rock": {
				Name:        "rock",
				ImageURL:    ptr("https://example.com/path/to/image/rock"),
				Description: ptr("rock glass description"),
			},
			"beer": {
				Name:        "beer",
				ImageURL:    ptr("https://example.com/path/to/image/beer"),
				Description: ptr("beer glass description"),
			},
		}

		recipeSaveLister := new(MockRecipeSaveLister)
		recipeSaveLister.On("AllGlassTypes").Return(glassTypes, nil)

		c := &copilot{recipe: recipeSaveLister}

		got, err := c.FindGlassType()
		assert.NoError(t, err)
		assert.Equal(t, want, got)
	})
}
