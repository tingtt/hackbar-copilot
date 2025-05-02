package order

import (
	"hackbar-copilot/internal/domain/order"
	"hackbar-copilot/internal/domain/user"
)

// SetUserInfo implements Order.
func (o *orderimpl) SetUserInfo(customerEmail order.CustomerEmail, customerName string, autofill bool) (user.User, error) {
	u := user.User{
		Email: user.Email(customerEmail),
		Name:  customerName,
	}
	if !autofill {
		u.NameConfirmed = true
	}
	return u, o.user.Save(u)
}
