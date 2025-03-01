package cashier

import (
	"hackbar-copilot/internal/domain/checkout"
	"time"
)

// LatestUnCachedOutCheckouts implements Cashier.
func (c *cashier) LatestUnCachedOutCheckouts() ([]checkout.Checkout, error) {
	var latestCashoutTimestamp *time.Time
	for cashout, err := range c.cashout.Latest() {
		if err != nil {
			return nil, err
		}
		latestCashoutTimestamp = &cashout.Timestamp
		break
	}

	checkouts := []checkout.Checkout{}
	for checkout, err := range c.checkout.Latest(checkout.Since(*latestCashoutTimestamp)) {
		if err != nil {
			return nil, err
		}
		if latestCashoutTimestamp != nil && checkout.Timestamp.Before(*latestCashoutTimestamp) {
			break
		}
		checkouts = append(checkouts, checkout)
	}
	return checkouts, nil
}
