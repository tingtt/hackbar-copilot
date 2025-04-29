package cashier

import (
	"errors"
	"fmt"
	"hackbar-copilot/internal/domain/checkout"
	"hackbar-copilot/internal/domain/order"
	usecaseutils "hackbar-copilot/internal/usecase/utils"
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

	for order_, err := range c.order.Latest(order.FilterCustomerEmail(customerEmail), order.IgnoreCheckedOut()) {
		if err != nil {
			return checkout.Checkout{}, err
		}
		if /* specified order */ orderIDsMap[order_.ID] {
			newCheckout.TotalPrice += order_.Price
			delete(orderIDsMap, order_.ID)
			_, err := updateOrderStatus(c.order, order_.ID, order.StatusCheckedOut, newCheckout.Timestamp)
			if err != nil {
				return checkout.Checkout{}, err
			}
		} else {
			_, err := updateOrderStatus(c.order, order_.ID, order.StatusCanceled, newCheckout.Timestamp)
			if err != nil {
				return checkout.Checkout{}, err
			}
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

func updateOrderStatus(order_ order.SaveFindListListener, id order.ID, status order.Status, timestamp time.Time) (order.Order, error) {
	o, err := order_.Find(id)
	if err != nil {
		if errors.Is(err, order.ErrNotFound) {
			return order.Order{}, usecaseutils.ErrNotFound
		}
		return order.Order{}, err
	}

	o.Status = status
	o.Timestamps = append(o.Timestamps, order.StatusUpdateTimestamp{
		Status:    status,
		Timestamp: timestamp,
	})

	err = order_.Save(o)
	if err != nil {
		return order.Order{}, err
	}
	return o, nil
}
