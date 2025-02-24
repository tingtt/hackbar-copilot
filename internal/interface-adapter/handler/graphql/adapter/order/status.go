package orderadapter

import (
	"hackbar-copilot/internal/domain/order"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
)

// Status implements OutputAdapter.
func (o *outputAdapter) Status(status order.Status) model.OrderStatus {
	switch status {
	case order.StatusOrdered:
		return model.OrderStatusOrdered
	case order.StatusPrepared:
		return model.OrderStatusPrepared
	case order.StatusDelivered:
		return model.OrderStatusDelivered
	case order.StatusCanceled:
		return model.OrderStatusCanceled
	case order.StatusCheckedOut:
		return model.OrderStatusCheckedout
	default:
		return model.OrderStatusUnknown
	}
}
