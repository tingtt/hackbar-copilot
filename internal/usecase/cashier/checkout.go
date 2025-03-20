package cashier

import (
	"fmt"
	"hackbar-copilot/internal/domain/checkout"
	"hackbar-copilot/internal/domain/order"
	"time"

	"github.com/google/uuid"
)

// Checkout implements Cashier.
func (c *cashier) Checkout(
	customerEmail order.CustomerEmail, orderIDs []order.ID, diffs []checkout.Diff, paymentType checkout.PaymentType,
) (checkout.Checkout, error) {
	newCheckout := checkout.Checkout{
		ID:            checkout.ID(uuid.NewString()),
		CustomerEmail: customerEmail,
		OrderIDs:      orderIDs,
		Diffs:         diffs,
		TotalPrice:    0,
		PaymentType:   paymentType,
		Timestamp:     time.Now(),
	}

	orderIDsMap := make(map[order.ID]bool, len(orderIDs))
	for _, id := range orderIDs {
		orderIDsMap[id] = true
	}

	for order, err := range c.order.Latest(order.IgnoreCheckedOut()) {
		if err != nil {
			return checkout.Checkout{}, err
		}
		if /* specified order */ orderIDsMap[order.ID] {
			newCheckout.TotalPrice += order.Price
			delete(orderIDsMap, order.ID)
		}
	}
	for notFoundOrderID := range orderIDsMap {
		return checkout.Checkout{}, fmt.Errorf("order not found or already checked out: %s", notFoundOrderID)
	}

	for _, diff := range diffs {
		newCheckout.TotalPrice += diff.Price
	}

	err := c.checkout.Save(newCheckout)
	if err != nil {
		return checkout.Checkout{}, err
	}
	return newCheckout, nil
}
