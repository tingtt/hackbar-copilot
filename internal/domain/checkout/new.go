package checkout

import (
	"hackbar-copilot/internal/domain/order"
	"time"

	"github.com/google/uuid"
)

func New(
	customerEmail order.CustomerEmail,
	orders []order.Order,
	diffs []Diff,
	paymentType PaymentType,
) (Checkout, error) {
	newCheckout := Checkout{
		ID:            ID(uuid.NewString()),
		CustomerEmail: customerEmail,
		Orders:        orders,
		Diffs:         diffs,
		TotalPrice:    totalPrice(orders, diffs),
		PaymentType:   paymentType,
		Timestamp:     time.Now().UTC(),
	}
	return newCheckout, newCheckout.Validate()
}

func totalPrice(orders []order.Order, diffs []Diff) float32 {
	totalPrice := float32(0)
	for _, order_ := range orders {
		if order_.Status != order.StatusCanceled {
			totalPrice += order_.Price
		}
	}
	for _, diff := range diffs {
		totalPrice += diff.Price
	}
	return totalPrice
}
