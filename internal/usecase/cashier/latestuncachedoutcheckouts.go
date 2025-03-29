package cashier

import (
	"hackbar-copilot/internal/domain/checkout"

	"github.com/tingtt/options"
)

// LatestUnCachedOutCheckouts implements Cashier.
func (c *cashier) LatestUnCachedOutCheckouts() ([]checkout.Checkout, error) {
	chechoutOptions := []options.Applier[checkout.ListerOption]{}
	for cashout, err := range c.cashout.Latest() {
		if err != nil {
			return nil, err
		}
		chechoutOptions = append(chechoutOptions, checkout.Since(cashout.Timestamp))
		break
	}

	checkouts := []checkout.Checkout{}
	for checkout, err := range c.checkout.Latest(chechoutOptions...) {
		if err != nil {
			return nil, err
		}
		checkouts = append(checkouts, checkout)
	}
	return checkouts, nil
}
