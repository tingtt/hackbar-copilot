package graphql

import (
	"testing"
)

func TestNewHandler(t *testing.T) {
	t.Parallel()

	t.Run("will return non-nil struct", func(t *testing.T) {
		t.Parallel()

		// TODO: Refactor to make it testable
		// handler := NewHandler(graph.Dependencies{}, Option{"test-secret"})

		// assert.NotNil(t, handler)
		// assert.NotNil(t, handler.ServeHTTP)
	})
}
