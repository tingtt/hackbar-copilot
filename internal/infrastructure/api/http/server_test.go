package http

import (
	"testing"
)

func TestNewServer(t *testing.T) {
	t.Parallel()

	t.Run("will return non-nil struct", func(t *testing.T) {
		t.Parallel()

		// TODO: Refactor to make it testable
		// server := NewServer("", graph.Dependencies{}, graphql.Option{JWTSecret: "test-secret"})

		// assert.NotNil(t, server)
		// assert.NotNil(t, server.Handler)
		// assert.NotNil(t, server.Handler.ServeHTTP)
	})
}
