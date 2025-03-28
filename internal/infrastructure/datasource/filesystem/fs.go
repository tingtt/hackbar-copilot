package filesystem

import (
	"hackbar-copilot/internal/domain/cashout"
	"hackbar-copilot/internal/domain/checkout"
	"hackbar-copilot/internal/domain/menu"
	"hackbar-copilot/internal/domain/order"
	"hackbar-copilot/internal/domain/recipe"
	"hackbar-copilot/internal/domain/stock"
	"hackbar-copilot/internal/domain/user"
)

type Filesystem interface {
	Recipe() recipe.Repository
	Menu() menu.Repository
	Stock() stock.Repository
	Order() (r order.Repository, close func())
	Cashout() cashout.Repository
	Checkout() checkout.Repository
	User() user.Repository
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

	fs := filesystem{read: fsR, write: fsW, data: data}

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
	read  fsR
	write fsW
	data  data
}

type data struct {
	users        []user.User
	recipeGroups []recipe.RecipeGroup
	recipeTypes  map[string]recipe.RecipeType
	glassTypes   map[string]recipe.GlassType
	menuItems    []menu.Item
	stocks       map[string]bool

	// orders
	//
	// orders is sorted by created desc.
	orders []order.Order

	// checkouts
	//
	// checkouts is sorted by created desc.
	checkouts []checkout.Checkout
}

func (data data) isEmpty() bool {
	return len(data.recipeGroups) == 0 &&
		len(data.recipeTypes) == 0 && len(data.glassTypes) == 0 &&
		len(data.menuItems) == 0 && len(data.orders) == 0
}

// Recipe implements Filesystem.
func (f *filesystem) Recipe() recipe.Repository {
	return &recipeRepository{f}
}

// Menu implements Filesystem.
func (f *filesystem) Menu() menu.Repository {
	return &menuRepository{f}
}

// Stock implements Filesystem.
func (f *filesystem) Stock() stock.Repository {
	return &stockRepository{f}
}

// Order implements Filesystem.
func (f *filesystem) Order() (_ order.Repository, _close func()) {
	r := &orderRepository{f, nil}
	return r, func() {
		if r.save != nil {
			close(r.save)
		}
	}
}

// Cashout implements Filesystem.
func (f *filesystem) Cashout() cashout.Repository {
	return &cashoutRepository{f}
}

// Checkout implements Filesystem.
func (f *filesystem) Checkout() checkout.Repository {
	return &checkoutRepository{f}
}

// User implements Filesystem.
func (f *filesystem) User() user.Repository {
	return newUserRepository(f)
}
