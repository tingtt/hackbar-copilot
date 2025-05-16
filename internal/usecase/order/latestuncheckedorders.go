package order

import (
	"hackbar-copilot/internal/domain/order"
)

// LatestUncheckedOrders implements Order.
func (o *orderimpl) LatestUncheckedOrders(customerEmail order.CustomerEmail) ([]order.Order, error) {
	orders := []order.Order{}
	for order, err := range o.datasource.Order().LatestUncheckedOrdersUser(customerEmail) {
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}
