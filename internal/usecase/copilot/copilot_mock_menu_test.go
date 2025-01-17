package copilot

import (
	"hackbar-copilot/internal/domain/menu"
	"iter"

	"github.com/stretchr/testify/mock"
)

var _ menu.SaveLister = new(MockMenuSaveLister)

type MockMenuSaveLister struct {
	mock.Mock
}

// All implements menu.SaveLister.
func (m *MockMenuSaveLister) All() iter.Seq2[menu.Group, error] {
	args := m.Called()
	return args.Get(0).(iter.Seq2[menu.Group, error])
}

// Save implements menu.SaveLister.
func (m *MockMenuSaveLister) Save(g menu.Group) error {
	args := m.Called(g)
	return args.Error(0)
}
