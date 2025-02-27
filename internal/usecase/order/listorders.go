package order

import (
	"hackbar-copilot/internal/domain/order"
)

// ListOrders implements Order.
func (o *orderimpl) ListOrders(customerID order.CustomerID) ([]order.Order, error) {
	orders := []order.Order{}
	for order, err := range o.order.Latest(
		order.FilterCustomerID(customerID), order.IgnoreCheckedOut(),
	) {
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}
