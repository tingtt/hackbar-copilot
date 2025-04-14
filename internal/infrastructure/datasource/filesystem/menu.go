package filesystem

import (
	"hackbar-copilot/internal/domain/menu"
	"iter"
	"slices"
)

var _ menu.Repository = (*menuRepository)(nil)

type menuRepository struct {
	fs *filesystem
}

// Find implements menu.Repository.
func (m *menuRepository) Find(itemName string, optionName string) (menu.ItemOption, error) {
	for _, mi := range m.fs.data.menuItems {
		if mi.Name == itemName {
			for _, mo := range mi.Options {
				if mo.Name == optionName {
					return mo, nil
				}
			}
		}
	}
	return menu.ItemOption{}, menu.ErrNotFound
}

// All implements menu.Repository.
func (m *menuRepository) All() iter.Seq2[menu.Item, error] {
	return func(yield func(menu.Item, error) bool) {
		for _, mi := range m.fs.data.menuItems {
			if !yield(mi, nil) {
				return
			}
		}
	}
}

// Save implements menu.Repository.
func (m *menuRepository) Save(g menu.Item) error {
	for i, savedMenuGroup := range m.fs.data.menuItems {
		if savedMenuGroup.Name == g.Name {
			m.fs.data.menuItems[i] = g
			return nil
		}
	}
	m.fs.data.menuItems = append(m.fs.data.menuItems, g)
	return nil
}

// Remove implements menu.Repository.
func (m *menuRepository) Remove(itemName string) error {
	for i, savedMenuGroup := range m.fs.data.menuItems {
		if savedMenuGroup.Name == itemName {
			m.fs.data.menuItems = slices.Delete(m.fs.data.menuItems, i, i+1)
			return nil
		}
	}
	return menu.ErrNotFound
}
