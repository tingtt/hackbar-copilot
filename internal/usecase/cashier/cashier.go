package cashier

import (
	"hackbar-copilot/internal/domain/cashout"
	"hackbar-copilot/internal/domain/checkout"
	"hackbar-copilot/internal/domain/order"
	"reflect"
	"time"
)

type Cashier interface {
	Checkout(
		customerID order.CustomerID, orderIDs []order.ID, diffs []checkout.Diff, paymentType checkout.PaymentType,
	) (checkout.Checkout, error)
	LatestUnCachedOutCheckouts() ([]checkout.Checkout, error)

	Cashout(staffID cashout.StaffID, checkoutIDs []checkout.ID) (cashout.Cashout, error)
	ListCashouts(since, until time.Time) ([]cashout.Cashout, error)
}

func New(deps Dependencies) Cashier {
	deps.validate()
	return &cashier{
		order:    order.NewSaveListListener(deps.Order),
		checkout: checkout.NewSaveLister(deps.Checkout),
		cashout:  cashout.NewRegisterLister(deps.Order, deps.Cashout),
	}
}

type Dependencies struct {
	Order    order.Repository
	Checkout checkout.Repository
	Cashout  cashout.Repository
}

func (d Dependencies) validate() {
	for i := range reflect.ValueOf(d).NumField() {
		if reflect.ValueOf(d).Field(i).IsNil() {
			t := reflect.TypeOf(d).Field(i).Type
			panic(t.PkgPath() + "." + t.Name() + " cannot be nil")
		}
	}
}

type cashier struct {
	order    order.Lister
	checkout checkout.SaveLister
	cashout  cashout.RegisterLister
}
