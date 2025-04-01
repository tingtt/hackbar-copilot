package menuadapter

import (
	"hackbar-copilot/internal/domain/menu"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
)

// MenuItems implements MenuAdapterOut.
func (m *outputAdapter) MenuItems(menuGroups []menu.Item, recipeGroups []*model.RecipeGroup) []*model.MenuItem {
	var groups []*model.MenuItem
	for _, menuGroup := range menuGroups {
		recipes := make(map[string]model.Recipe)
		for _, recipeGroup := range recipeGroups {
			if recipeGroup.Name == menuGroup.Name {
				for _, recipe := range recipeGroup.Recipes {
					recipes[recipe.Name] = *recipe
				}
			}
		}
		groups = append(groups, m.menuItem(menuGroup, recipes))
	}
	return groups
}

func (m *outputAdapter) menuItem(menuGroup menu.Item, recipes map[string]model.Recipe) *model.MenuItem {
	minPrice, options := m.menuItemOptions(menuGroup.Options, recipes)
	return &model.MenuItem{
		Name:        menuGroup.Name,
		ImageURL:    menuGroup.ImageURL,
		Flavor:      menuGroup.Flavor,
		Options:     options,
		MinPriceYen: float64(minPrice),
	}
}

func (m *outputAdapter) menuItemOptions(menuItems []menu.ItemOption, recipes map[string]model.Recipe) (minPrice float64, options []*model.MenuItemOption) {
	for _, menuItem := range menuItems {
		if minPrice == 0 {
			minPrice = float64(menuItem.Price)
		}
		if float64(menuItem.Price) < minPrice {
			minPrice = float64(menuItem.Price)
		}
		recipe, ok := recipes[menuItem.Name]
		if ok {
			options = append(options, m.menuItemOption(menuItem, &recipe))
		} else {
			options = append(options, m.menuItemOption(menuItem, nil))
		}
	}
	return minPrice, options
}

func (m *outputAdapter) menuItemOption(menuItem menu.ItemOption, recipe *model.Recipe) *model.MenuItemOption {
	return &model.MenuItemOption{
		Name:       menuItem.Name,
		Category:   menuItem.Category,
		ImageURL:   menuItem.ImageURL,
		Materials:  menuItem.Materials,
		OutOfStock: menuItem.OutOfStock,
		PriceYen:   float64(menuItem.Price),
		Recipe:     recipe,
	}
}
