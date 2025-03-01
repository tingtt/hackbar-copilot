package cashout

import (
	"hackbar-copilot/internal/domain/checkout"
	"time"
)

// Register implements RegisterLister.
func (s *registerLister) Register(staffID StaffID, checkouts []checkout.Checkout) (Cashout, error) {
	summary := Cashout{
		Checkouts: checkouts,
		Revenue:   0,
		Timestamp: time.Now(),
		StaffID:   staffID,
	}
	for _, c := range checkouts {
		summary.Revenue += c.TotalPrice
	}

	err := s.Repository.Save(summary)
	if err != nil {
		return Cashout{}, err
	}
	return summary, nil
}
