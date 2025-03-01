package copilot

import (
	"hackbar-copilot/internal/domain/cashout"
	"iter"

	"github.com/stretchr/testify/mock"
	"github.com/tingtt/options"
)

var _ cashout.Repository = new(MockCashout)

type MockCashout struct {
	mock.Mock
}

// Latest implements cashout.Repository.
func (m *MockCashout) Latest(optionAppliers ...options.Applier[cashout.ListerOption]) iter.Seq2[cashout.Cashout, error] {
	args := m.Called(optionAppliers)
	return args.Get(0).(iter.Seq2[cashout.Cashout, error])
}

// Save implements cashout.Repository.
func (m *MockCashout) Save(s cashout.Cashout) error {
	args := m.Called(s)
	return args.Error(0)
}
