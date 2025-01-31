package copilot

import (
	"hackbar-copilot/internal/domain/menu"
	"hackbar-copilot/internal/domain/recipe"
	"hackbar-copilot/internal/domain/stock"
	"hackbar-copilot/internal/usecase/sort"
)

type Copilot interface {
	ListRecipes(sortFunc sort.Yield[recipe.RecipeGroup]) ([]recipe.RecipeGroup, error)
	SaveRecipe(rg recipe.RecipeGroup) error
	SaveRecipeType(rt recipe.RecipeType) error
	SaveGlassType(gt recipe.GlassType) error
	FindRecipeGroup(name string) (recipe.RecipeGroup, error)
	FindRecipeType() (map[string]recipe.RecipeType, error)
	FindGlassType() (map[string]recipe.GlassType, error)

	SaveAsMenuGroup(recipeGroupName string, arg SaveAsMenuGroupArg) (menu.Group, error)
	ListMenu(sortFunc sort.Yield[menu.Group]) ([]menu.Group, error)

	Materials(sortFunc sort.Yield[stock.Material], optionAppliers ...QueryOptionApplier) ([]stock.Material, error)
	UpdateStock(inStockMaterials, outOfStockMaterials []string) error
}

type SaveAsMenuGroupArg struct {
	Flavor *string
	Items  map[string]MenuFromRecipeGroupArg
}

type MenuFromRecipeGroupArg struct {
	ImageURL *string
	Price    int
}

func New(deps Dependencies) Copilot {
	return &copilot{
		recipe: recipe.NewSaveLister(deps.Recipe),
		menu:   menu.NewSaveLister(deps.Menu),
		stock:  stock.NewSaveLister(deps.Stock),
	}
}

type Dependencies struct {
	Recipe recipe.Repository
	Menu   menu.Repository
	Stock  stock.Repository
}

type copilot struct {
	recipe recipe.SaveLister
	menu   menu.SaveLister
	stock  stock.SaveLister
}
