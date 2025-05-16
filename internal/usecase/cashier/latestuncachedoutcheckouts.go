package cashier

import (
	"hackbar-copilot/internal/domain/checkout"
)

// LatestUnCachedOutCheckouts implements Cashier.
func (c *cashier) LatestUnCachedOutCheckouts() ([]checkout.Checkout, error) {
	checkouts := []checkout.Checkout{}
	for checkout, err := range c.datasource.Checkout().LatestUnCachedOutCheckouts() {
		if err != nil {
			return nil, err
		}
		checkouts = append(checkouts, checkout)
	}
	return checkouts, nil
}
