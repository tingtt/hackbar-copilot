package copilot

import (
	"hackbar-copilot/internal/domain/order"

	"github.com/tingtt/options"
)

// LatestOrders implements Copilot.
func (c *copilot) LatestOrders() ([]order.Order, error) {
	optionAppliers := []options.Applier[order.ListerOption]{}
	for summary, err := range c.ordersummary.Latest() {
		if err != nil {
			return nil, err
		}
		optionAppliers = append(optionAppliers, order.Since(summary.Timestamp))
		break
	}

	orders := []order.Order{}
	for order, err := range c.order.Latest(optionAppliers...) {
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}
