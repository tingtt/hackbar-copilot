package adapter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMenuAdapterOut(t *testing.T) {
	t.Parallel()

	t.Run("will return non-nil struct", func(t *testing.T) {
		t.Parallel()

		adapter := NewMenuAdapterOut()

		assert.NotNil(t, adapter)
	})
}
