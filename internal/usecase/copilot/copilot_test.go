package copilot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()

	t.Run("will panic on not enough dependencies", func(t *testing.T) {
		t.Parallel()

		assert.Panics(t, func() {
			New(Dependencies{})
		})
	})

	t.Run("will return non-nil struct", func(t *testing.T) {
		t.Parallel()

		copilot := New(Dependencies{new(MockGateway)}).(*copilot)

		assert.NotNil(t, copilot)
		assert.NotNil(t, copilot.datasource)
	})
}
