package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewResolver(t *testing.T) {
	t.Parallel()

	t.Run("will return non-nil struct", func(t *testing.T) {
		t.Parallel()

		resolver := NewResolver(Dependencies{})

		assert.NotNil(t, resolver)
	})

	t.Run("will insert local dependency", func(t *testing.T) {
		t.Parallel()

		resolver := NewResolver(Dependencies{}).(*Resolver)

		assert.NotNil(t, resolver.recipeAdapter)
	})
}
