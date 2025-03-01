package cashout

import (
	"hackbar-copilot/internal/domain/checkout"
	"time"
)

type Cashout struct {
	Checkouts []checkout.Checkout
	Revenue   float32
	Timestamp time.Time
	StaffID   StaffID
}

type StaffID string
