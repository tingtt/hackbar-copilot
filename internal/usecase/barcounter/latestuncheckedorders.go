package barcounter

import (
	"hackbar-copilot/internal/domain/order"
)

// LatestUncheckedOrders implements BarCounter.
func (c *barcounterimpl) LatestUncheckedOrders() ([]order.Order, error) {
	orders := []order.Order{}
	for order, err := range c.datasource.Order().LatestUncheckedOrders() {
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}
