package ordertest

import (
	"hackbar-copilot/internal/domain/order"
	"iter"
)

func IterWithNilError(items []order.Order) iter.Seq2[order.Order, error] {
	return func(yield func(order.Order, error) bool) {
		for _, item := range items {
			if !yield(item, nil) {
				break
			}
		}
	}
}
