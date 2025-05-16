package filesystem

import (
	"hackbar-copilot/internal/domain/menu"
	"hackbar-copilot/internal/usecase/copilot"
	orderusecase "hackbar-copilot/internal/usecase/order"
	usecaseutils "hackbar-copilot/internal/usecase/utils"
	"iter"
	"slices"
	"sync"
)

var _ copilot.MenuSaveListRemover = (*menuRepository)(nil)
var _ orderusecase.MenuFindLister = (*menuRepository)(nil)

type menuRepository struct {
	fs    *filesystem
	mutex *sync.RWMutex
}

// Find implements menu.Repository.
func (r *menuRepository) Find(itemName string, optionName string) (menu.ItemOption, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, mi := range r.fs.data.menuItems {
		if mi.Name == itemName {
			for _, mo := range mi.Options {
				if mo.Name == optionName {
					return mo, nil
				}
			}
		}
	}
	return menu.ItemOption{}, usecaseutils.ErrNotFound
}

// All implements menu.Repository.
func (r *menuRepository) All() iter.Seq2[menu.Item, error] {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	return func(yield func(menu.Item, error) bool) {
		for _, mi := range r.fs.data.menuItems {
			if !yield(mi, nil) {
				return
			}
		}
	}
}

// Save implements menu.Repository.
func (r *menuRepository) Save(g menu.Item) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i, savedMenuGroup := range r.fs.data.menuItems {
		if savedMenuGroup.Name == g.Name {
			r.fs.data.menuItems[i] = g
			return nil
		}
	}
	r.fs.data.menuItems = append(r.fs.data.menuItems, g)
	return nil
}

// Remove implements menu.Repository.
func (r *menuRepository) Remove(itemName string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i, savedMenuGroup := range r.fs.data.menuItems {
		if savedMenuGroup.Name == itemName {
			r.fs.data.menuItems = slices.Delete(r.fs.data.menuItems, i, i+1)
			return nil
		}
	}
	return usecaseutils.ErrNotFound
}
