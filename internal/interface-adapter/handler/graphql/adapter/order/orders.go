package orderadapter

import (
	"hackbar-copilot/internal/domain/order"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
)

// Orders implements OutputAdapter.
func (o *outputAdapter) Orders(orders []order.Order) []*model.Order {
	convertedOrders := []*model.Order{}
	for _, order := range orders {
		convertedOrders = append(convertedOrders, o.Order(order))
	}
	return convertedOrders
}
