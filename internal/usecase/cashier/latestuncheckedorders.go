package cashier

import (
	"hackbar-copilot/internal/domain/order"
)

// LatestUncheckedOrders implements Cashier.
func (c *cashier) LatestUncheckedOrders() ([]order.Order, error) {
	orders := []order.Order{}
	for order_, err := range c.datasource.Order().LatestUncheckedOrders() {
		if err != nil {
			return nil, err
		}
		orders = append(orders, order_)
	}
	return orders, nil
}
