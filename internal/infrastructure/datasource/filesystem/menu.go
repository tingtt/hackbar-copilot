package filesystem

import (
	"hackbar-copilot/internal/domain/menu"
	"iter"
)

var _ menu.Repository = (*menuRepository)(nil)

type menuRepository struct {
	fs *filesystem
}

// All implements menu.Repository.
func (m *menuRepository) All() iter.Seq2[menu.Group, error] {
	return func(yield func(menu.Group, error) bool) {
		for _, mg := range m.fs.data.menuGroups {
			if !yield(mg, nil) {
				return
			}
		}
	}
}

// Save implements menu.Repository.
func (m *menuRepository) Save(g menu.Group) error {
	for i, savedMenuGroup := range m.fs.data.menuGroups {
		if savedMenuGroup.Name == g.Name {
			m.fs.data.menuGroups[i] = g
			return nil
		}
	}
	m.fs.data.menuGroups = append(m.fs.data.menuGroups, g)
	return nil
}
