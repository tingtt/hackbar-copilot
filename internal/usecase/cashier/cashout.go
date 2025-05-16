package cashier

import (
	"fmt"
	"hackbar-copilot/internal/domain/cashout"
	"hackbar-copilot/internal/domain/checkout"
	"slices"
)

// Cashout implements Cashier.
func (c *cashier) Cashout(staffID cashout.StaffID, checkoutIDs []checkout.ID) (cashout.Cashout, error) {
	checkouts := []checkout.Checkout{}
	for checkout, err := range c.datasource.Checkout().LatestUnCachedOutCheckouts() {
		if err != nil {
			return cashout.Cashout{}, err
		}
		if slices.Contains(checkoutIDs, checkout.ID) {
			checkouts = append(checkouts, checkout)
		}
	}

	newCashout, err := cashout.New(staffID, checkouts)
	if err != nil {
		return cashout.Cashout{}, fmt.Errorf("failed to create new cashout: %w", err)
	}

	err = c.datasource.Cashout().Save(newCashout)
	if err != nil {
		return cashout.Cashout{}, fmt.Errorf("failed to save new cashout: %w", err)
	}

	return newCashout, nil
}
