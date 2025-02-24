package copilot

import (
	"hackbar-copilot/internal/domain/menu"
	"iter"

	"github.com/stretchr/testify/mock"
)

var _ menu.SaveFindLister = new(MockMenu)

type MockMenu struct {
	mock.Mock
}

// Find implements menu.SaveFindLister.
func (m *MockMenu) Find(groupName string, itemName string) (menu.Item, error) {
	args := m.Called(groupName, itemName)
	return args.Get(0).(menu.Item), args.Error(1)
}

// All implements menu.SaveLister.
func (m *MockMenu) All() iter.Seq2[menu.Group, error] {
	args := m.Called()
	return args.Get(0).(iter.Seq2[menu.Group, error])
}

// Save implements menu.SaveLister.
func (m *MockMenu) Save(g menu.Group) error {
	args := m.Called(g)
	return args.Error(0)
}
