package cashier

import (
	"fmt"
	"hackbar-copilot/internal/domain/checkout"
	"hackbar-copilot/internal/domain/order"
	"hackbar-copilot/internal/utils/sliceutil"
	"iter"
	"time"
)

// Checkout implements Cashier.
func (c *cashier) Checkout(
	customerEmail order.CustomerEmail,
	orderIDs []order.ID,
	diffs []checkout.Diff,
	paymentType checkout.PaymentType,
) (checkout.Checkout, error) {
	currentTimestamp := time.Now().UTC()

	orders := make([]order.Order, 0, len(orderIDs))
	for order_, err := range iterateOrders(c.datasource.Order(), orderIDs, customerEmail) {
		if err != nil {
			return checkout.Checkout{}, err
		}
		if order_.specified {
			orders = append(orders,
				order_.ApplyStatus(order.StatusCheckedOut, currentTimestamp),
			)
		} else {
			orders = append(orders,
				order_.ApplyStatus(order.StatusCanceled, currentTimestamp),
			)
		}
	}

	newCheckout, err := checkout.New(customerEmail, orders, diffs, paymentType)
	if err != nil {
		return checkout.Checkout{}, fmt.Errorf("failed to create new checkout: %w", err)
	}

	err = c.datasource.Checkout().Save(newCheckout)
	if err != nil {
		return checkout.Checkout{}, fmt.Errorf("failed to save new checkout: %w", err)
	}
	err = c.datasource.Order().Remove(sliceutil.Map(orders,
		func(o order.Order) order.ID { return o.ID },
	)...)
	if err != nil {
		return checkout.Checkout{}, fmt.Errorf("failed to remove orders: %w", err)
	}

	return newCheckout, nil
}

type iterableOrder struct {
	order.Order
	specified bool
}

// iterateOrders iterates over the orders and yields them.
func iterateOrders(orderLister OrderLister, orderIDs []order.ID, customerEmail order.CustomerEmail) iter.Seq2[iterableOrder, error] {
	return func(yield func(iterableOrder, error) bool) {
		orderIDsMap := make(map[order.ID]bool, len(orderIDs))
		for _, id := range orderIDs {
			orderIDsMap[id] = true
		}

		for order_, err := range orderLister.LatestUncheckedOrders() {
			if err != nil {
				yield(iterableOrder{}, err)
				return
			}
			if /* specified order */ orderIDsMap[order_.ID] {
				yield(iterableOrder{Order: order_, specified: true}, nil)
			} else if order_.CustomerEmail == customerEmail {
				yield(iterableOrder{Order: order_, specified: false}, nil)
			}
		}
	}
}
