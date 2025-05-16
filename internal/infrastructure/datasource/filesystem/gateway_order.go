package filesystem

import (
	orderusecase "hackbar-copilot/internal/usecase/order"
)

var _ orderusecase.Gateway = (*orderGateway)(nil)

type orderGateway struct {
	*gateway
}

// Menu implements order.Gateway.
func (o *orderGateway) Menu() orderusecase.MenuFindLister {
	return &o.menu
}

// Order implements order.Gateway.
func (o *orderGateway) Order() orderusecase.OrderSaveLister {
	return &o.order
}

// User implements order.Gateway.
func (o *orderGateway) User() orderusecase.UserSaveGetter {
	return &o.user
}
