package recipeadapter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRecipeAdapter(t *testing.T) {
	t.Parallel()

	t.Run("will return non-nil struct", func(t *testing.T) {
		t.Parallel()

		adapter := New()

		assert.NotNil(t, adapter)
	})

	t.Run("will return non-nil fields", func(t *testing.T) {
		t.Parallel()

		adapter := New().(*recipeAdapter)

		assert.NotNil(t, adapter.InputAdapter)
		assert.NotNil(t, adapter.OutputAdapter)
	})
}
