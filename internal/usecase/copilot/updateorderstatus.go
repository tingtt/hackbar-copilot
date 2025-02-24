package copilot

import (
	"errors"
	"hackbar-copilot/internal/domain/order"
	usecaseutils "hackbar-copilot/internal/usecase/utils"
	"time"
)

// UpdateOrderStatus implements Copilot.
func (c *copilot) UpdateOrderStatus(id order.ID, status order.Status, timestamp time.Time) (order.Order, error) {
	o, err := c.order.Find(id)
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

	err = c.order.Save(o)
	if err != nil {
		return order.Order{}, err
	}
	return o, nil
}
