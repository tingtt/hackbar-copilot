package cashier

import (
	"hackbar-copilot/internal/domain/cashout"
	"time"
)

// ListCashouts implements Cashier.
func (c *cashier) ListCashouts(since, until time.Time) ([]cashout.Cashout, error) {
	cashouts := []cashout.Cashout{}
	for cashout_, err := range c.cashout.Latest(cashout.Since(since), cashout.Until(until)) {
		if err != nil {
			return nil, err
		}
		cashouts = append(cashouts, cashout_)
	}
	return cashouts, nil
}
