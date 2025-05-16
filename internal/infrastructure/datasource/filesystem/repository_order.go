package filesystem

import (
	"hackbar-copilot/internal/domain/order"
	"hackbar-copilot/internal/usecase/barcounter"
	"hackbar-copilot/internal/usecase/cashier"
	usecaseutils "hackbar-copilot/internal/usecase/utils"
	"iter"
	"sync"
)

var _ barcounter.OrderSaveListFinder = (*orderRepository)(nil)
var _ cashier.OrderListRemover = (*orderRepository)(nil)

type orderRepository struct {
	fs    *filesystem
	mutex *sync.RWMutex
}

// Find implements barcounter.OrderSaveListFinder.
func (r *orderRepository) Find(id order.ID) (order.Order, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, o := range r.fs.data.uncheckedOrders {
		if o.ID == id {
			return o, nil
		}
	}
	return order.Order{}, usecaseutils.ErrNotFound
}

// LatestUncheckedOrders implements barcounter.OrderSaveListFinder.
func (r *orderRepository) LatestUncheckedOrders() iter.Seq2[order.Order, error] {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	return func(yield func(order.Order, error) bool) {
		for _, o := range r.fs.data.uncheckedOrders {
			if !yield(o, nil) {
				return
			}
		}
	}
}

// LatestUncheckedOrdersUser implements order.OrderSaveLister.
func (r *orderRepository) LatestUncheckedOrdersUser(customerEmail order.CustomerEmail) iter.Seq2[order.Order, error] {
	return func(yield func(order.Order, error) bool) {
		for order_, err := range r.LatestUncheckedOrders() {
			if err != nil {
				if !yield(order.Order{}, err) {
					return
				}
				return
			}
			if order_.CustomerEmail == customerEmail {
				if !yield(order_, nil) {
					return
				}
			}
		}
	}
}

// Save implements barcounter.OrderSaveListFinder.
func (r *orderRepository) Save(d order.Order) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i, savedOrder := range r.fs.data.uncheckedOrders {
		if savedOrder.ID == d.ID {
			r.fs.data.uncheckedOrders[i] = d
			return nil
		}
	}
	r.fs.data.uncheckedOrders = append([]order.Order{d}, r.fs.data.uncheckedOrders...)
	return nil
}

// Remove implements cashier.OrderListRemover.
func (r *orderRepository) Remove(ids ...order.ID) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for _, id := range ids {
		for i, savedOrder := range r.fs.data.uncheckedOrders {
			if savedOrder.ID == id {
				r.fs.data.uncheckedOrders = append(
					r.fs.data.uncheckedOrders[:i],
					r.fs.data.uncheckedOrders[i+1:]...,
				)
				break
			}
		}
	}
	return nil
}
