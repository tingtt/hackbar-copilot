package copilot

import (
	"hackbar-copilot/internal/domain/stock"
	"iter"

	"github.com/stretchr/testify/mock"
)

var _ StockSaveLister = new(MockStock)

type MockStock struct {
	mock.Mock
}

// All implements stock.SaveLister.
func (m *MockStock) All() iter.Seq2[stock.Material, error] {
	args := m.Called()
	return args.Get(0).(iter.Seq2[stock.Material, error])
}

// Save implements stock.SaveLister.
func (m *MockStock) Save(inStockMaterials []string, outOfStockMaterials []string) error {
	args := m.Called(inStockMaterials, outOfStockMaterials)
	return args.Error(0)
}
