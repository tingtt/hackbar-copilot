package orderadapter

import (
	"hackbar-copilot/internal/domain/order"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
)

// Order implements OutputAdapter.
func (o *outputAdapter) Order(order order.Order) *model.Order {
	return &model.Order{
		ID:            string(order.ID),
		CustomerEmail: string(order.CustomerEmail),
		CustomerName:  order.CustomerName,
		MenuID: &model.MenuID{
			ItemName:   order.MenuItemID.ItemName,
			OptionName: order.MenuItemID.OptionName,
		},
		Timestamps: o.StatusUpdateTimestamps(order.Timestamps),
		Status:     o.Status(order.Status),
		Price:      float64(order.Price),
	}
}
