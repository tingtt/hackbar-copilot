package graphql

import (
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHandler(t *testing.T) {
	t.Parallel()

	t.Run("will return non-nil struct", func(t *testing.T) {
		t.Parallel()

		handler := NewHandler(graph.Dependencies{})

		assert.NotNil(t, handler)
		assert.NotNil(t, handler.ServeHTTP)
	})
}
