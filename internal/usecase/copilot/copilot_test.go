package copilot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()

	t.Run("will return non-nil struct", func(t *testing.T) {
		t.Parallel()

		copilot := New(Dependencies{
			Recipe: new(MockRecipeSaveLister),
			Menu:   new(MockMenuSaveLister),
			Stock:  new(MockStockSaveLister),
		}).(*copilot)

		assert.NotNil(t, copilot)
		assert.NotNil(t, copilot.recipe)
		assert.NotNil(t, copilot.menu)
		assert.NotNil(t, copilot.stock)
	})
}
