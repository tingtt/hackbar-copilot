package copilot

import (
	"hackbar-copilot/internal/domain/order"

	"github.com/tingtt/options"
)

// LatestOrders implements Copilot.
func (c *copilot) LatestOrders() ([]order.Order, error) {
	optionAppliers := []options.Applier[order.ListerOption]{}
	for cashout, err := range c.cashout.Latest() {
		if err != nil {
			return nil, err
		}
		// apply filter since by latest cashout timestamp
		optionAppliers = append(optionAppliers, order.Since(cashout.Timestamp))
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
