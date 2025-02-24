package menuadapter

import (
	"hackbar-copilot/internal/domain/menu"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
)

// MenuGroups implements MenuAdapterOut.
func (m *outputAdapter) MenuGroups(menuGroups []menu.Group, recipeGroups []*model.RecipeGroup) []*model.MenuGroup {
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

func (m *outputAdapter) menuGroup(menuGroup menu.Group, recipes map[string]model.Recipe) *model.MenuGroup {
	minPrice, items := m.menuItems(menuGroup.Items, recipes)
	return &model.MenuGroup{
		Name:        menuGroup.Name,
		ImageURL:    menuGroup.ImageURL,
		Flavor:      menuGroup.Flavor,
		Items:       items,
		MinPriceYen: float64(minPrice),
	}
}

func (m *outputAdapter) menuItems(menuItems []menu.Item, recipes map[string]model.Recipe) (minPrice float64, items []*model.MenuItem) {
	for _, menuItem := range menuItems {
		if minPrice == 0 {
			minPrice = float64(menuItem.Price)
		}
		if float64(menuItem.Price) < minPrice {
			minPrice = float64(menuItem.Price)
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

func (m *outputAdapter) menuItem(menuItem menu.Item, recipe *model.Recipe) *model.MenuItem {
	return &model.MenuItem{
		Name:       menuItem.Name,
		ImageURL:   menuItem.ImageURL,
		Materials:  menuItem.Materials,
		OutOfStock: menuItem.OutOfStock,
		PriceYen:   float64(menuItem.Price),
		Recipe:     recipe,
	}
}
