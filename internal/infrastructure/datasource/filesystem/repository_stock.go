package filesystem

import (
	"hackbar-copilot/internal/domain/stock"
	"hackbar-copilot/internal/usecase/copilot"
	"iter"
	"sync"
)

var _ copilot.StockSaveLister = (*stockRepository)(nil)

type stockRepository struct {
	fs    *filesystem
	mutex *sync.RWMutex
}

// All implements stock.Repository.
func (r *stockRepository) All() iter.Seq2[stock.Material, error] {
	return func(yield func(stock.Material, error) bool) {
		r.mutex.RLock()
		defer r.mutex.RUnlock()

		for name, inStock := range r.fs.data.stocks {
			if !yield(stock.Material{Name: name, InStock: inStock}, nil) {
				return
			}
		}
	}
}

// Save implements stock.Repository.
func (r *stockRepository) Save(inStockMaterials, outOfStockMaterials []string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.fs.data.stocks == nil {
		r.fs.data.stocks = make(map[string]bool)
	}
	for _, name := range inStockMaterials {
		r.fs.data.stocks[name] = true
	}
	for _, name := range outOfStockMaterials {
		r.fs.data.stocks[name] = false
	}
	return nil
}
