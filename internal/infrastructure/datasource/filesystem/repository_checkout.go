package filesystem

import (
	"hackbar-copilot/internal/domain/checkout"
	"hackbar-copilot/internal/usecase/cashier"
	"iter"
	"slices"
	"sync"
)

var _ cashier.CheckoutSaveListRemover = (*checkoutRepository)(nil)

type checkoutRepository struct {
	fs    *filesystem
	mutex *sync.RWMutex
}

// Save implements cashier.CheckoutSaveListRemover.
func (r *checkoutRepository) Save(o checkout.Checkout) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i, savedCheckout := range r.fs.data.uncashedoutCheckouts {
		if savedCheckout.ID == o.ID {
			r.fs.data.uncashedoutCheckouts[i] = o
			return nil
		}
	}
	r.fs.data.uncashedoutCheckouts = append([]checkout.Checkout{o}, r.fs.data.uncashedoutCheckouts...)
	return nil
}

// LatestUnCachedOutCheckouts implements cashier.CheckoutSaveListRemover.
func (r *checkoutRepository) LatestUnCachedOutCheckouts() iter.Seq2[checkout.Checkout, error] {
	return func(yield func(checkout.Checkout, error) bool) {
		r.mutex.RLock()
		defer r.mutex.RUnlock()

		for _, checkout := range r.fs.data.uncashedoutCheckouts {
			if !yield(checkout, nil) {
				return
			}
		}
	}
}

// Remove implements cashier.CheckoutSaveListRemover.
func (r *checkoutRepository) Remove(ids ...checkout.ID) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for _, id := range ids {
		for i, checkout := range r.fs.data.uncashedoutCheckouts {
			if checkout.ID == id {
				r.fs.data.uncashedoutCheckouts = slices.Delete(r.fs.data.uncashedoutCheckouts, i, i+1)
				break
			}
		}
	}
	return nil
}
