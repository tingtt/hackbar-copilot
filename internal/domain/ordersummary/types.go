package ordersummary

import (
	"hackbar-copilot/internal/domain/order"
	"time"
)

type Summary struct {
	Orders    []order.Order
	Revenue   float32
	Timestamp time.Time
}
