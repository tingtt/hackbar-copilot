package filesystem

import (
	"hackbar-copilot/internal/domain/stock"
	"iter"
)

var _ stock.Repository = (*stockRepository)(nil)

type stockRepository struct {
	fs *filesystem
}

// All implements stock.Repository.
func (s *stockRepository) All() iter.Seq2[stock.Material, error] {
	return func(yield func(stock.Material, error) bool) {
		for name, inStock := range s.fs.data.stocks {
			if !yield(stock.Material{Name: name, InStock: inStock}, nil) {
				return
			}
		}
	}
}

// Save implements stock.Repository.
func (s *stockRepository) Save(inStockMaterials, outOfStockMaterials []string) error {
	if s.fs.data.stocks == nil {
		s.fs.data.stocks = make(map[string]bool)
	}
	for _, name := range inStockMaterials {
		s.fs.data.stocks[name] = true
	}
	for _, name := range outOfStockMaterials {
		s.fs.data.stocks[name] = false
	}
	return nil
}
