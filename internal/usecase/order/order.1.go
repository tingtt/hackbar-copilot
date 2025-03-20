package order

import (
	"errors"
	"hackbar-copilot/internal/domain/order"
	"hackbar-copilot/internal/domain/user"
	"time"

	"github.com/google/uuid"
)

// Order implements Order.
func (o *orderimpl) Order(customerEmail order.CustomerEmail, customerName *string, menuItemID order.MenuItemID) (order.Order, error) {
	if customerName == nil {
		u, err := o.user.Get(user.Email(customerEmail))
		if err != nil {
			if errors.Is(err, user.ErrNotFound) {
				return order.Order{}, errors.New("customer name not specified")
			}
			return order.Order{}, err
		}
		customerName = &u.Name
	}

	menuItem, err := o.menu.Find(menuItemID.ItemName, menuItemID.OptionName)
	if err != nil {
		return order.Order{}, err
	}

	new := order.Order{
		ID:            order.ID(uuid.NewString()),
		CustomerEmail: customerEmail,
		CustomerName:  *customerName,
		MenuItemID:    menuItemID,
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
