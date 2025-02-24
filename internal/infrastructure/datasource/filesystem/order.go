package filesystem

import (
	"hackbar-copilot/internal/domain/order"
	"iter"

	"github.com/tingtt/options"
)

var _ order.Repository = (*orderRepository)(nil)

type orderRepository struct {
	fs   *filesystem
	save chan order.SavedOrder
}

// Find implements order.Repository.
func (o *orderRepository) Find(id order.ID) (order.Order, error) {
	for _, o := range o.fs.data.orders {
		if o.ID == id {
			return o, nil
		}
	}
	return order.Order{}, order.ErrNotFound
}

// Latest implements order.Repository.
func (o *orderRepository) Latest(optionAppliers ...options.Applier[order.ListerOption]) iter.Seq2[order.Order, error] {
	option := options.Create(optionAppliers...)

	return func(yield func(order.Order, error) bool) {
		for _, o := range o.fs.data.orders {
			if option.Since != nil && o.CreatedAt().Before(*option.Since) {
				continue
			}
			if option.CustomerID != nil && o.CustomerID != *option.CustomerID {
				continue
			}
			if !yield(o, nil) {
				return
			}
		}
	}
}

// Listen implements order.Repository.
func (o *orderRepository) Listen() (chan order.SavedOrder, error) {
	if o.save == nil {
		o.save = make(chan order.SavedOrder)
	}
	return o.save, nil
}

// Save implements order.Repository.
func (o *orderRepository) Save(newOrder order.Order) error {
	for i, currentOrder := range o.fs.data.orders {
		if currentOrder.ID == newOrder.ID {
			o.fs.data.orders[i] = newOrder
			o.notify(order.SavedOrder{Order: newOrder})
			return nil
		}
	}
	o.fs.data.orders = append([]order.Order{newOrder}, o.fs.data.orders...)
	o.notify(order.SavedOrder{Order: newOrder})
	return nil
}

func (o *orderRepository) notify(s order.SavedOrder) {
	if o.save != nil {
		o.save <- s
	}
}
