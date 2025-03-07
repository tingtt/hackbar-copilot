package order

import (
	"hackbar-copilot/internal/domain/menu"
	"hackbar-copilot/internal/domain/order"
	"hackbar-copilot/internal/usecase/sort"
	"reflect"
)

type Order interface {
	ListMenu(sortFunc sort.Yield[menu.Item]) ([]menu.Item, error)
	Order(customerID order.CustomerID, menuItemID order.MenuItemID) (order.Order, error)
	ListUncheckedOrders(customerID order.CustomerID) ([]order.Order, error)
}

func New(deps Dependencies) Order {
	deps.validate()
	return &orderimpl{
		menu:  menu.NewFindLister(deps.Menu),
		order: order.NewSaveListListener(deps.Order),
	}
}

type Dependencies struct {
	Menu  menu.Repository
	Order order.Repository
}

func (d Dependencies) validate() {
	for i := range reflect.ValueOf(d).NumField() {
		if reflect.ValueOf(d).Field(i).IsNil() {
			t := reflect.TypeOf(d).Field(i).Type
			panic(t.PkgPath() + "." + t.Name() + " cannot be nil")
		}
	}
}

type orderimpl struct {
	menu  menu.FindLister
	order order.SaveFindListListener
}
