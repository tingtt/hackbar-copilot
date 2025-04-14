package copilot

import (
	"hackbar-copilot/internal/domain/menu"
	"iter"

	"github.com/stretchr/testify/mock"
)

var _ menu.SaveFindListRemover = new(MockMenu)

type MockMenu struct {
	mock.Mock
}

// Find implements menu.SaveFindLister.
func (m *MockMenu) Find(itemName string, optionName string) (menu.ItemOption, error) {
	args := m.Called(itemName, optionName)
	return args.Get(0).(menu.ItemOption), args.Error(1)
}

// All implements menu.SaveLister.
func (m *MockMenu) All() iter.Seq2[menu.Item, error] {
	args := m.Called()
	return args.Get(0).(iter.Seq2[menu.Item, error])
}

// Save implements menu.SaveLister.
func (m *MockMenu) Save(g menu.Item) error {
	args := m.Called(g)
	return args.Error(0)
}

// Remove implements menu.SaveFindListRemover.
func (m *MockMenu) Remove(itemName string) error {
	args := m.Called(itemName)
	return args.Error(0)
}
