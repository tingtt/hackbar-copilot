package cashout

import (
	"hackbar-copilot/internal/domain/checkout"
	"time"
)

func New(staffID StaffID, checkouts []checkout.Checkout) (Cashout, error) {
	summary := Cashout{
		Checkouts: checkouts,
		Revenue:   0,
		Timestamp: time.Now(),
		StaffID:   staffID,
	}
	for _, c := range checkouts {
		summary.Revenue += c.TotalPrice
	}
	return summary, nil
}
