package order

import (
	"hackbar-copilot/internal/domain/order"
	"time"

	"github.com/google/uuid"
)

// Order implements Order.
func (o *orderimpl) Order(customerID order.CustomerID, menuItemID order.MenuItemID) (order.Order, error) {
	menuItem, err := o.menu.Find(menuItemID.GroupName, menuItemID.ItemName)
	if err != nil {
		return order.Order{}, err
	}

	new := order.Order{
		ID:         order.ID(uuid.NewString()),
		CustomerID: customerID,
		MenuItemID: menuItemID,
		Timestamps: []order.StatusUpdateTimestamp{
			{
				Status:    order.StatusOrdered,
				Timestamp: time.Now(),
			},
		},
		Status: order.StatusOrdered,
		Price:  menuItem.Price,
	}
	return new, o.order.Save(new)
}
