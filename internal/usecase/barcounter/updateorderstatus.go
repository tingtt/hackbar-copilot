package barcounter

import (
	"hackbar-copilot/internal/domain/order"
	"time"
)

// UpdateOrderStatus implements BarCounter.
func (c *barcounterimpl) UpdateOrderStatus(id order.ID, status order.Status, timestamp time.Time) (order.Order, error) {
	o, err := c.datasource.Order().Find(id)
	if err != nil {
		return order.Order{}, err
	}

	o.Status = status
	o.Timestamps = append(o.Timestamps, order.StatusUpdateTimestamp{
		Status:    status,
		Timestamp: timestamp,
	})

	err = c.datasource.Order().Save(o)
	if err != nil {
		return order.Order{}, err
	}
	return o, nil
}
