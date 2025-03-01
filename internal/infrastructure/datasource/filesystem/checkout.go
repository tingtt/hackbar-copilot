package filesystem

import (
	"hackbar-copilot/internal/domain/checkout"
	"iter"

	"github.com/tingtt/options"
)

var _ checkout.Repository = (*checkoutRepository)(nil)

type checkoutRepository struct {
	fs *filesystem
}

// Latest implements checkout.Repository.
func (c *checkoutRepository) Latest(optionAppliers ...options.Applier[checkout.ListerOption]) iter.Seq2[checkout.Checkout, error] {
	option := options.Create(optionAppliers...)

	return func(yield func(checkout.Checkout, error) bool) {
		for _, checkout := range c.fs.data.checkouts {
			if option.Since != nil && checkout.Timestamp.Before(*option.Since) {
				break
			}
			if option.CustomerID != nil && checkout.CustomerID != *option.CustomerID {
				continue
			}
			if !yield(checkout, nil) {
				return
			}
		}
	}
}

// Save implements checkout.Repository.
func (c *checkoutRepository) Save(o checkout.Checkout) error {
	for i, savedCheckout := range c.fs.data.checkouts {
		if savedCheckout.ID == o.ID {
			c.fs.data.checkouts[i] = o
			return nil
		}
	}
	c.fs.data.checkouts = append([]checkout.Checkout{o}, c.fs.data.checkouts...)
	return nil
}
