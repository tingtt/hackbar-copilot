package barcounter

import (
	"hackbar-copilot/internal/domain/order"
	"iter"
	"reflect"
	"time"
)

type BarCounter interface {
	LatestUncheckedOrders() ([]order.Order, error)
	UpdateOrderStatus(id order.ID, status order.Status, timestamp time.Time) (order.Order, error)
}

func New(deps Dependencies) BarCounter {
	deps.validate()
	return &barcounterimpl{deps.Gateway}
}

type Dependencies struct {
	Gateway Gateway
}

func (d Dependencies) validate() {
	for i := range reflect.ValueOf(d).NumField() {
		if reflect.ValueOf(d).Field(i).IsNil() {
			t := reflect.TypeOf(d).Field(i).Type
			panic(t.PkgPath() + "." + t.Name() + " cannot be nil")
		}
	}
}

type barcounterimpl struct {
	datasource Gateway
}

type Gateway interface {
	Order() OrderSaveListFinder
}

type OrderSaveListFinder interface {
	Save(d order.Order) error
	LatestUncheckedOrders() iter.Seq2[order.Order, error]
	Find(id order.ID) (order.Order, error)
}
