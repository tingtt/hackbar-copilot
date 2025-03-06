package copilot

import (
	"hackbar-copilot/internal/domain/cashout"
	"hackbar-copilot/internal/domain/menu"
	"hackbar-copilot/internal/domain/order"
	"hackbar-copilot/internal/domain/recipe"
	"hackbar-copilot/internal/domain/stock"
	"hackbar-copilot/internal/usecase/sort"
	"reflect"
	"time"
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

	LatestUncheckedOrders() ([]order.Order, error)
	ListenOrder() (chan order.SavedOrder, error)
	UpdateOrderStatus(id order.ID, status order.Status, timestamp time.Time) (order.Order, error)
}

type SaveAsMenuGroupArg struct {
	Flavor *string
	Items  map[string]MenuFromRecipeGroupArg
}

type MenuFromRecipeGroupArg struct {
	ImageURL *string
	Price    float32
}

func New(deps Dependencies) Copilot {
	deps.validate()
	return &copilot{
		recipe:  recipe.NewSaveLister(deps.Recipe),
		menu:    menu.NewSaveLister(deps.Menu),
		stock:   stock.NewSaveLister(deps.Stock),
		order:   order.NewSaveListListener(deps.Order),
		cashout: cashout.NewRegisterLister(deps.Order, deps.Cashout),
	}
}

type Dependencies struct {
	Recipe  recipe.Repository
	Menu    menu.Repository
	Stock   stock.Repository
	Order   order.Repository
	Cashout cashout.Repository
}

func (d Dependencies) validate() {
	for i := range reflect.ValueOf(d).NumField() {
		if reflect.ValueOf(d).Field(i).IsNil() {
			t := reflect.TypeOf(d).Field(i).Type
			panic(t.PkgPath() + "." + t.Name() + " cannot be nil")
		}
	}
}

type copilot struct {
	recipe  recipe.SaveLister
	menu    menu.SaveFindLister
	stock   stock.SaveLister
	order   order.SaveFindListListener
	cashout cashout.Lister
}
