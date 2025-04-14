package recipe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSaveLister(t *testing.T) {
	t.Parallel()

	t.Run("will return non-nil struct", func(t *testing.T) {
		t.Parallel()

		s := NewSaveLister(new(MockRepository)).(*saverLister)
		assert.NotNil(t, s)
		assert.NotNil(t, s.Repository)
	})
}
