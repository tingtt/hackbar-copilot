package order

import (
	"fmt"
	"hackbar-copilot/internal/domain/order"
	"hackbar-copilot/internal/domain/user"
)

// SetUserInfo implements Order.
func (o *orderimpl) SetUserInfo(customerEmail order.CustomerEmail, customerName string, autofill bool) (user.User, error) {
	u, err := user.New(user.Email(customerEmail), customerName, !autofill)
	if err != nil {
		return user.User{}, fmt.Errorf("failed to create new user: %w", err)
	}

	err = o.datasource.User().Save(u)
	if err != nil {
		return user.User{}, fmt.Errorf("failed to save user: %w", err)
	}

	return u, nil
}
