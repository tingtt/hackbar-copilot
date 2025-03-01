package copilot

import (
	"hackbar-copilot/internal/domain/order"
	"iter"

	"github.com/stretchr/testify/mock"
	"github.com/tingtt/options"
)

var _ order.Repository = new(MockOrder)

type MockOrder struct {
	mock.Mock
}

// Find implements order.Repository.
func (m *MockOrder) Find(id order.ID) (order.Order, error) {
	args := m.Called(id)
	return args.Get(0).(order.Order), args.Error(1)
}

// Latest implements order.Repository.
func (m *MockOrder) Latest(optionAppliers ...options.Applier[order.ListerOption]) iter.Seq2[order.Order, error] {
	args := m.Called(optionAppliers)
	return args.Get(0).(iter.Seq2[order.Order, error])
}

// Listen implements order.Repository.
func (m *MockOrder) Listen() (chan order.SavedOrder, error) {
	args := m.Called()
	return args.Get(0).(chan order.SavedOrder), args.Error(1)
}

// Save implements order.Repository.
func (m *MockOrder) Save(o order.Order) error {
	args := m.Called(o)
	return args.Error(0)
}
