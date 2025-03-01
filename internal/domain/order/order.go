package order

import (
	"errors"
	"iter"
	"time"

	"github.com/tingtt/options"
)

type SaveFindListListener interface {
	saver
	finder
	Lister
	listener
}

type saver interface {
	Save(o Order) error
}

type ListerOption struct {
	Since            *time.Time
	CustomerID       *CustomerID
	IgnoreCheckedOut bool
}

func Since(t time.Time) options.Applier[ListerOption] {
	return func(lo *ListerOption) {
		lo.Since = &t
	}
}

func FilterCustomerID(id CustomerID) options.Applier[ListerOption] {
	return func(lo *ListerOption) {
		lo.CustomerID = &id
	}
}

func IgnoreCheckedOut() options.Applier[ListerOption] {
	return func(lo *ListerOption) {
		lo.IgnoreCheckedOut = true
	}
}

type Lister interface {
	Latest(optionAppliers ...options.Applier[ListerOption]) iter.Seq2[Order, error]
}

type listener interface {
	Listen() (chan SavedOrder, error)
}

var ErrNotFound = errors.New("order not found")

type finder interface {
	Find(id ID) (Order, error)
}

type Repository SaveFindListListener

func NewSaveListListener(r Repository) SaveFindListListener {
	return &saveListListener{r}
}

type saveListListener struct {
	Repository
}
