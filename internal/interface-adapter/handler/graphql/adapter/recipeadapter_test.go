package adapter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRecipeAdapter(t *testing.T) {
	t.Parallel()

	t.Run("will return non-nil struct", func(t *testing.T) {
		t.Parallel()

		adapter := NewRecipeAdapter()

		assert.NotNil(t, adapter)
	})

	t.Run("will return non-nil fields", func(t *testing.T) {
		t.Parallel()

		adapter := NewRecipeAdapter().(*recipeAdapter)

		assert.NotNil(t, adapter.RecipeAdapterIn)
		assert.NotNil(t, adapter.RecipeAdapterOut)
	})
}
