package copilot

import (
	"hackbar-copilot/internal/domain/menu"
	"hackbar-copilot/internal/domain/recipe"
	"hackbar-copilot/internal/domain/stock"
	"hackbar-copilot/internal/usecase/sort"
	"iter"
	"reflect"
)

type Copilot interface {
	ListRecipes(sortFunc sort.Yield[recipe.RecipeGroup]) ([]recipe.RecipeGroup, error)
	SaveRecipe(rg recipe.RecipeGroup) error
	RemoveRecipeAndMenuItem(name string) error
	SaveRecipeType(rt recipe.RecipeType) error
	SaveGlassType(gt recipe.GlassType) error
	FindRecipeGroup(name string) (recipe.RecipeGroup, error)
	FindRecipeType() (map[string]recipe.RecipeType, error)
	FindGlassType() (map[string]recipe.GlassType, error)

	SaveAsMenuItem(recipeGroupName string, arg SaveAsMenuItemArg) (menu.Item, error)

	Materials(sortFunc sort.Yield[stock.Material], optionAppliers ...QueryOptionApplier) ([]stock.Material, error)
	UpdateStock(inStockMaterials, outOfStockMaterials []string) error
}

type SaveAsMenuItemArg struct {
	Flavor  *string
	Options map[string]MenuFromRecipeGroupArg
	Remove  bool
}

type MenuFromRecipeGroupArg struct {
	ImageURL *string
	Price    float32
}

func New(deps Dependencies) Copilot {
	deps.validate()
	return &copilot{deps.Gateway}
}

type Dependencies struct {
	Gateway Gateway
}

func (d Dependencies) validate() {
	for i := range reflect.ValueOf(d).NumField() {
		if reflect.ValueOf(d).Field(i).IsNil() {
			t := reflect.TypeFor[Dependencies]().Field(i).Type
			panic(t.PkgPath() + "." + t.Name() + " cannot be nil")
		}
	}
}

type copilot struct {
	datasource Gateway
}

type Gateway interface {
	Recipe() RecipeSaveListRemover
	Menu() MenuSaveListRemover
	Stock() StockSaveLister
}

type RecipeSaveListRemover interface {
	Save(rg recipe.RecipeGroup) error
	SaveRecipeType(rt recipe.RecipeType) error
	SaveGlassType(gt recipe.GlassType) error
	All() iter.Seq2[recipe.RecipeGroup, error]
	AllRecipeTypes() iter.Seq2[recipe.RecipeType, error]
	AllGlassTypes() iter.Seq2[recipe.GlassType, error]
	Remove(recipeGroupName string) error
}

type MenuSaveListRemover interface {
	Save(g menu.Item) error
	All() iter.Seq2[menu.Item, error]
	Remove(itemName string) error
}

type StockSaveLister interface {
	Save(inStockMaterials, outOfStockMaterials []string) error
	All() iter.Seq2[stock.Material, error]
}
