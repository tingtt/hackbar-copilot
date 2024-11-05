package http

import (
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	t.Parallel()

	t.Run("will return non-nil struct", func(t *testing.T) {
		t.Parallel()

		server := NewServer("", graph.Dependencies{})

		assert.NotNil(t, server)
		assert.NotNil(t, server.Handler)
		assert.NotNil(t, server.Handler.ServeHTTP)
	})
}
