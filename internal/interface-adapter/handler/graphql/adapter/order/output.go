package orderadapter

import (
	"hackbar-copilot/internal/domain/order"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
)

type OutputAdapter interface {
	Orders(orders []order.Order) []*model.Order
	Order(order order.Order) *model.Order

	Status(status order.Status) model.OrderStatus

	StatusUpdateTimestamps(statusUpdates []order.StatusUpdateTimestamp) []*model.OrderStatusUpdateTimestamp
	StatusUpdateTimestamp(statusUpdate order.StatusUpdateTimestamp) *model.OrderStatusUpdateTimestamp
}

func NewOutputAdapter() OutputAdapter {
	return &outputAdapter{}
}

type outputAdapter struct{}
