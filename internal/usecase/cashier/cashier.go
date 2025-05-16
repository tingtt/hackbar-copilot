package cashier

import (
	"hackbar-copilot/internal/domain/cashout"
	"hackbar-copilot/internal/domain/checkout"
	"hackbar-copilot/internal/domain/order"
	"iter"
	"reflect"
	"time"

	"github.com/tingtt/options"
)

type Cashier interface {
	LatestUncheckedOrders() ([]order.Order, error)

	// Checkout saves a checkout and removes related orders from temporary storage.
	Checkout(
		customerEmail order.CustomerEmail, orderIDs []order.ID, diffs []checkout.Diff, paymentType checkout.PaymentType,
	) (checkout.Checkout, error)
	LatestUnCachedOutCheckouts() ([]checkout.Checkout, error)

	// Cashout saves a cashout and removes related checkouts from temporary storage.
	Cashout(staffID cashout.StaffID, checkoutIDs []checkout.ID) (cashout.Cashout, error)
	ListCashouts(since, until time.Time) ([]cashout.Cashout, error)
}

func New(deps Dependencies) Cashier {
	deps.validate()
	return &cashier{deps.Gateway}
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

type cashier struct {
	datasource Gateway
}

type Gateway interface {
	Order() OrderListRemover
	Checkout() CheckoutSaveListRemover
	Cashout() CashoutSaveLister
}

type OrderListRemover interface {
	Remove(ids ...order.ID) error
	OrderLister
}

type OrderLister interface {
	LatestUncheckedOrders() iter.Seq2[order.Order, error]
}

type CheckoutSaveListRemover interface {
	Save(o checkout.Checkout) error
	Remove(ids ...checkout.ID) error
	LatestUnCachedOutCheckouts() iter.Seq2[checkout.Checkout, error]
}

type CashoutSaveLister interface {
	Save(cashout cashout.Cashout) error
	Latest(optionAppliers ...options.Applier[ListerOption]) iter.Seq2[cashout.Cashout, error]
}

type ListerOption struct {
	Since *time.Time
	Until *time.Time
}

func Since(t time.Time) options.Applier[ListerOption] {
	return func(lo *ListerOption) {
		lo.Since = &t
	}
}

func Until(t time.Time) options.Applier[ListerOption] {
	return func(lo *ListerOption) {
		lo.Until = &t
	}
}
