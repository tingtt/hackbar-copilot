package orderadapter

import (
	"fmt"
	"hackbar-copilot/internal/domain/order"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
)

// ApplyStatus implements InputAdapter.
func (i *inputAdapter) ApplyStatus(status model.OrderStatus) (order.Status, error) {
	switch status {
	case model.OrderStatusOrdered:
		return order.StatusOrdered, nil
	case model.OrderStatusPrepared:
		return order.StatusPrepared, nil
	case model.OrderStatusDelivered:
		return order.StatusDelivered, nil
	case model.OrderStatusCanceled:
		return order.StatusCanceled, nil
	case model.OrderStatusCheckedout:
		return order.StatusCheckedOut, nil
	default:
		return order.Status(""), fmt.Errorf("invalid status: %s", status)
	}
}
