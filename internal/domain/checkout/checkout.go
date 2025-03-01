package checkout

import (
	"hackbar-copilot/internal/domain/order"
	"iter"
	"time"

	"github.com/tingtt/options"
)

type SaveLister interface {
	saver
	Lister
}

type saver interface {
	Save(o Checkout) error
}

type ListerOption struct {
	Since      *time.Time
	CustomerID *order.CustomerID
}

func Since(t time.Time) options.Applier[ListerOption] {
	return func(lo *ListerOption) {
		lo.Since = &t
	}
}

func FilterCustomerID(id order.CustomerID) options.Applier[ListerOption] {
	return func(lo *ListerOption) {
		lo.CustomerID = &id
	}
}

type Lister interface {
	Latest(optionAppliers ...options.Applier[ListerOption]) iter.Seq2[Checkout, error]
}

type Repository SaveLister

func NewSaveLister(r Repository) SaveLister {
	return &saveLister{r}
}

func NewLister(r Repository) Lister {
	return &saveLister{r}
}

type saveLister struct {
	Repository
}
