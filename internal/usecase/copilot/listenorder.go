package copilot

import (
	"hackbar-copilot/internal/domain/order"
)

// ListenOrder implements Copilot.
func (c *copilot) ListenOrder() (chan order.SavedOrder, error) {
	return c.order.Listen()
}
