package order

import (
	"errors"
	"fmt"
	"hackbar-copilot/internal/domain/order"
	"hackbar-copilot/internal/domain/user"
)

// Order implements Order.
func (o *orderimpl) Order(customerEmail order.CustomerEmail, customerName *string, menuItemID order.MenuItemID) (order.Order, error) {
	if customerName == nil {
		u, err := o.datasource.User().Get(user.Email(customerEmail))
		if err != nil {
			return order.Order{}, err
		}
		if u.Name == "" {
			return order.Order{}, errors.New("customer name not specified")
		}
		customerName = &u.Name
	} else {
		_, err := o.SetUserInfo(customerEmail, *customerName, false /* not autofill */)
		if err != nil {
			return order.Order{}, err
		}
	}

	menuItem, err := o.datasource.Menu().Find(menuItemID.ItemName, menuItemID.OptionName)
	if err != nil {
		return order.Order{}, err
	}

	newOrder, err := order.New(
		order.Customer{
			Email: customerEmail,
			Name:  *customerName,
		},
		order.MenuItem{
			ID:    menuItemID,
			Price: menuItem.Price,
		},
	)
	if err != nil {
		return order.Order{}, fmt.Errorf("failed to create new order: %w", err)
	}

	err = o.datasource.Order().Save(newOrder)
	if err != nil {
		return order.Order{}, fmt.Errorf("failed to save new order: %w", err)
	}

	return newOrder, nil
}
