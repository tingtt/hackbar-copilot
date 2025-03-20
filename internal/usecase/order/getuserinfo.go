package order

import (
	"hackbar-copilot/internal/domain/order"
	"hackbar-copilot/internal/domain/user"
)

// GetUserInfo implements Order.
func (o *orderimpl) GetUserInfo(customerEmail order.CustomerEmail) (user.User, error) {
	return o.user.Get(user.Email(customerEmail))
}
