package cashier

import (
	"hackbar-copilot/internal/domain/cashout"
	"hackbar-copilot/internal/domain/checkout"
	"slices"

	"github.com/tingtt/options"
)

// Cashout implements Cashier.
func (c *cashier) Cashout(staffID cashout.StaffID, checkoutIDs []checkout.ID) (cashout.Cashout, error) {
	optionAppliers := []options.Applier[checkout.ListerOption]{}
	for cashout_, err := range c.cashout.Latest() {
		if err != nil {
			return cashout.Cashout{}, err
		}
		optionAppliers = append(optionAppliers, checkout.Since(cashout_.Timestamp))
		break
	}

	checkouts := []checkout.Checkout{}
	for checkout, err := range c.checkout.Latest(optionAppliers...) {
		if err != nil {
			return cashout.Cashout{}, err
		}
		// TODO: skipped order never referenced
		if slices.Contains(checkoutIDs, checkout.ID) {
			checkouts = append(checkouts, checkout)
		}
	}

	return c.cashout.Register(staffID, checkouts)
}
