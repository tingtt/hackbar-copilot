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

		copilot := New(Dependencies{
			Recipe:  new(MockRecipeSaveLister),
			Menu:    new(MockMenu),
			Stock:   new(MockStockSaveLister),
			Order:   new(MockOrder),
			Cashout: new(MockCashout),
		}).(*copilot)

		assert.NotNil(t, copilot)
		assert.NotNil(t, copilot.recipe)
		assert.NotNil(t, copilot.menu)
		assert.NotNil(t, copilot.stock)
		assert.NotNil(t, copilot.order)
		assert.NotNil(t, copilot.cashout)
	})
}
