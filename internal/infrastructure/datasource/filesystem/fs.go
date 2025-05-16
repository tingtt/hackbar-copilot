package filesystem

import (
	"hackbar-copilot/internal/domain/checkout"
	"hackbar-copilot/internal/domain/menu"
	"hackbar-copilot/internal/domain/order"
	"hackbar-copilot/internal/domain/recipe"
	"hackbar-copilot/internal/domain/user"
	"hackbar-copilot/internal/usecase/barcounter"
	"hackbar-copilot/internal/usecase/cashier"
	"hackbar-copilot/internal/usecase/copilot"
	orderusecase "hackbar-copilot/internal/usecase/order"
)

type Filesystem interface {
	BarCounterGateway() barcounter.Gateway
	CashierGateway() cashier.Gateway
	CopilotGateway() copilot.Gateway
	OrderGateway() orderusecase.Gateway
	SavePersistently() error
}

// NewRepository creates a new instance that implements
// filesystem.Filesystem and recipes.Repository.
//
// Example usage:
//
//	repository := filesystem.NewRepository("/path/to/data-dir")
//	err := repository.Save(...)
//	// handle error ...
//	err = repository.SavePersistently()
//	// handle error ...
//
// It create files to baseDir follow the structure:
//
//	```
//	/path/to/data-dir
//	├── 0_user.toml									// inmemory loaded
//	├── 1_recipe_groups.toml				// inmemory loaded
//	├── 2_recipe_types.toml					// inmemory loaded
//	├── 3_glass_types.toml					// inmemory loaded
//	├── 4_menu_items.toml					  // inmemory loaded
//	├── 5_stocks.toml								// inmemory loaded
//	├── 6_orders.toml								// inmemory loaded
//	├── 7_checkouts.toml						// inmemory loaded
//	└── 8_cashout_<timestamp>.toml
//	```
func NewRepository(baseDir string) (Filesystem, error) {
	fsR := newFSR(baseDir)
	fsW := newFSW(baseDir)
	data, err := loadData(fsR)
	if err != nil {
		return nil, err
	}

	fs := filesystem{fsR, fsW, data, gateway{}}
	fs.initializeGateway()

	if fs.data.isEmpty() {
		// check base dir is writable
		err = fs.SavePersistently()
		if err != nil {
			return nil, err
		}
	}

	return &fs, nil
}

type filesystem struct {
	read    fsR
	write   fsW
	data    data
	gateway gateway
}

type data struct {
	users        []user.User
	recipeGroups []recipe.RecipeGroup
	recipeTypes  map[string]recipe.RecipeType
	glassTypes   map[string]recipe.GlassType
	menuItems    []menu.Item
	stocks       map[string]bool

	// uncheckedOrders
	//
	// uncheckedOrders is sorted by created desc.
	uncheckedOrders []order.Order

	// uncashedoutCheckouts
	//
	// uncashedoutCheckouts is sorted by created desc.
	uncashedoutCheckouts []checkout.Checkout
}

func (data data) isEmpty() bool {
	return len(data.recipeGroups) == 0 &&
		len(data.recipeTypes) == 0 && len(data.glassTypes) == 0 &&
		len(data.menuItems) == 0 && len(data.uncheckedOrders) == 0
}
