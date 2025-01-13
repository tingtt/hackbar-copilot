package adapter

import (
	"hackbar-copilot/internal/domain/menu"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
)

// MenuGroups implements MenuAdapterOut.
func (m *menuAdapterOut) MenuGroups(menuGroups []menu.Group, recipeGroups []*model.RecipeGroup) []*model.MenuGroup {
	var groups []*model.MenuGroup
	for _, menuGroup := range menuGroups {
		recipes := make(map[string]model.Recipe)
		for _, recipeGroup := range recipeGroups {
			if recipeGroup.Name == menuGroup.Name {
				for _, recipe := range recipeGroup.Recipes {
					recipes[recipe.Name] = *recipe
				}
			}
		}
		groups = append(groups, m.menuGroup(menuGroup, recipes))
	}
	return groups
}

func (m *menuAdapterOut) menuGroup(menuGroup menu.Group, recipes map[string]model.Recipe) *model.MenuGroup {
	minPrice, items := m.menuItems(menuGroup.Items, recipes)
	return &model.MenuGroup{
		Name:        menuGroup.Name,
		ImageURL:    menuGroup.ImageURL,
		Flavor:      menuGroup.Flavor,
		Items:       items,
		MinPriceYen: minPrice,
	}
}

func (m *menuAdapterOut) menuItems(menuItems []menu.Item, recipes map[string]model.Recipe) (minPrice int, items []*model.MenuItem) {
	for _, menuItem := range menuItems {
		if minPrice == 0 {
			minPrice = menuItem.Price
		}
		if menuItem.Price < minPrice {
			minPrice = menuItem.Price
		}
		recipe, ok := recipes[menuItem.Name]
		if ok {
			items = append(items, m.menuItem(menuItem, &recipe))
		} else {
			items = append(items, m.menuItem(menuItem, nil))
		}
	}
	return minPrice, items
}

func (m *menuAdapterOut) menuItem(menuItem menu.Item, recipe *model.Recipe) *model.MenuItem {
	return &model.MenuItem{
		Name:       menuItem.Name,
		ImageURL:   menuItem.ImageURL,
		Materials:  menuItem.Materials,
		OutOfStock: menuItem.OutOfStock,
		PriceYen:   menuItem.Price,
		Recipe:     recipe,
	}
}
