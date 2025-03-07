package menutest

import (
	"hackbar-copilot/internal/domain/menu"
	"iter"
)

func IterWithNilError(items []menu.Item) iter.Seq2[menu.Item, error] {
	return func(yield func(menu.Item, error) bool) {
		for _, item := range DeepCopy(items) {
			if !yield(item, nil) {
				break
			}
		}
	}
}

func DeepCopy(groups []menu.Item) []menu.Item {
	copied := make([]menu.Item, len(groups))
	for i, group := range groups {
		copied[i] = DeepCopyGroup(group)
	}
	return copied
}

func DeepCopyGroup(group menu.Item) menu.Item {
	return menu.Item{
		Name:     group.Name,
		ImageURL: group.ImageURL,
		Flavor:   group.Flavor,
		Options:  DeepCopyItems(group.Options),
	}
}

func DeepCopyItems(items []menu.ItemOption) []menu.ItemOption {
	copied := make([]menu.ItemOption, len(items))
	for i, item := range items {
		copied[i] = DeepCopyItem(item)
	}
	return copied
}

func DeepCopyItem(item menu.ItemOption) menu.ItemOption {
	MaterialsCopy := make([]string, len(item.Materials))
	copy(MaterialsCopy, item.Materials)

	return menu.ItemOption{
		Name:       item.Name,
		ImageURL:   item.ImageURL,
		Materials:  MaterialsCopy,
		OutOfStock: item.OutOfStock,
		Price:      item.Price,
	}
}
