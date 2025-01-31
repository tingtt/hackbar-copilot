package menutest

import (
	"hackbar-copilot/internal/domain/menu"
	"iter"
)

func ptr[T any](v T) *T {
	return &v
}

func IterWithNilError(items []menu.Group) iter.Seq2[menu.Group, error] {
	return func(yield func(menu.Group, error) bool) {
		for _, item := range DeepCopy(items) {
			if !yield(item, nil) {
				break
			}
		}
	}
}

func DeepCopy(groups []menu.Group) []menu.Group {
	copied := make([]menu.Group, len(groups))
	for i, group := range groups {
		copied[i] = DeepCopyGroup(group)
	}
	return copied
}

func DeepCopyGroup(group menu.Group) menu.Group {
	return menu.Group{
		Name:     group.Name,
		ImageURL: group.ImageURL,
		Flavor:   group.Flavor,
		Items:    DeepCopyItems(group.Items),
	}
}

func DeepCopyItems(items []menu.Item) []menu.Item {
	copied := make([]menu.Item, len(items))
	for i, item := range items {
		copied[i] = DeepCopyItem(item)
	}
	return copied
}

func DeepCopyItem(item menu.Item) menu.Item {
	MaterialsCopy := make([]string, len(item.Materials))
	copy(MaterialsCopy, item.Materials)

	return menu.Item{
		Name:       item.Name,
		ImageURL:   item.ImageURL,
		Materials:  MaterialsCopy,
		OutOfStock: item.OutOfStock,
		Price:      item.Price,
	}
}
