package order

import (
	"hackbar-copilot/internal/domain/menu"
	"hackbar-copilot/internal/domain/order"
	"hackbar-copilot/internal/usecase/sort"
)

type Order interface {
	ListMenu(sortFunc sort.Yield[menu.Group]) ([]menu.Group, error)
	Order(customerID order.CustomerID, menuItemID order.MenuItemID) (order.Order, error)
	ListOrders(customerID order.CustomerID) ([]order.Order, error)
}

func New(deps Dependencies) Order {
	return &orderimpl{
		menu:  menu.NewFindLister(deps.Menu),
		order: order.NewSaveListListener(deps.Order),
	}
}

type Dependencies struct {
	Menu  menu.Repository
	Order order.Repository
}

type orderimpl struct {
	menu  menu.FindLister
	order order.SaveFindListListener
}
