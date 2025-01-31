package stocktest

import (
	"hackbar-copilot/internal/domain/stock"
	"iter"
)

func IterWithNilError(items []stock.Material) iter.Seq2[stock.Material, error] {
	return func(yield func(stock.Material, error) bool) {
		for _, item := range items {
			if !yield(item, nil) {
				break
			}
		}
	}
}
