package cashout

import (
	"hackbar-copilot/internal/domain/checkout"
	"hackbar-copilot/internal/domain/order"
	"iter"
	"time"

	"github.com/tingtt/options"
)

type registerer interface {
	Register(staffID StaffID, checkouts []checkout.Checkout) (Cashout, error)
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

type Lister interface {
	Latest(optionAppliers ...options.Applier[ListerOption]) iter.Seq2[Cashout, error]
}

type RegisterLister interface {
	registerer
	Lister
}

type Repository interface {
	Lister
	Save(s Cashout) error
}

func NewRegisterLister(or order.Repository, r Repository) RegisterLister {
	return &registerLister{or, r}
}

type registerLister struct {
	orderRepository order.Repository
	Repository
}
