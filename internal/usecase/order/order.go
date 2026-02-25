package order

import (
	"hackbar-copilot/internal/domain/menu"
	"hackbar-copilot/internal/domain/order"
	"hackbar-copilot/internal/domain/user"
	"hackbar-copilot/internal/usecase/sort"
	"iter"
	"reflect"
)

type Order interface {
	GetUserInfo(customerEmail order.CustomerEmail) (user.User, error)
	SetUserInfo(customerEmail order.CustomerEmail, customerName string, autofill bool) (user.User, error)
	ListMenu(sortFunc sort.Yield[menu.Item]) ([]menu.Item, error)
	Order(customerEmail order.CustomerEmail, customerName *string, menuItemID order.MenuItemID) (order.Order, error)
	LatestUncheckedOrders(customerEmail order.CustomerEmail) ([]order.Order, error)
}

func New(deps Dependencies) Order {
	deps.validate()
	return &orderimpl{deps.Gateway}
}

type Dependencies struct {
	Gateway Gateway
}

func (d Dependencies) validate() {
	for i := range reflect.ValueOf(d).NumField() {
		if reflect.ValueOf(d).Field(i).IsNil() {
			t := reflect.TypeFor[Dependencies]().Field(i).Type
			panic(t.PkgPath() + "." + t.Name() + " cannot be nil")
		}
	}
}

type orderimpl struct {
	datasource Gateway
}

type Gateway interface {
	User() UserSaveGetter
	Menu() MenuFindLister
	Order() OrderSaveLister
}

type UserSaveGetter interface {
	Save(d user.User) error
	Get(email user.Email) (user.User, error)
}

type MenuFindLister interface {
	Find(itemName, optionName string) (menu.ItemOption, error)
	All() iter.Seq2[menu.Item, error]
}

type OrderSaveLister interface {
	Save(d order.Order) error
	LatestUncheckedOrdersUser(customerEmail order.CustomerEmail) iter.Seq2[order.Order, error]
}
