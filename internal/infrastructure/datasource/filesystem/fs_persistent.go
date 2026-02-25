package filesystem

import (
	"errors"
	"hackbar-copilot/internal/infrastructure/datasource/filesystem/toml"
	"os"
)

func loadData(fs fsR) (d data, err error) {
	err = loadFromToml(fs, "0_user.toml", "user", &d.users)
	if err != nil {
		return data{}, err
	}
	err = loadFromToml(fs, "1_recipe_groups.toml", "recipe_group", &d.recipeGroups)
	if err != nil {
		return data{}, err
	}
	err = loadFromToml(fs, "2_recipe_types.toml", "recipe_type", &d.recipeTypes)
	if err != nil {
		return data{}, err
	}
	err = loadFromToml(fs, "3_glass_types.toml", "glass_type", &d.glassTypes)
	if err != nil {
		return data{}, err
	}
	err = loadFromToml(fs, "4_menu_items.toml", "menu_items", &d.menuItems)
	if err != nil {
		return data{}, err
	}
	err = loadFromToml(fs, "5_stocks.toml", "stock", &d.stocks)
	if err != nil {
		return data{}, err
	}
	err = loadFromToml(fs, "6_orders.toml", "order", &d.uncheckedOrders)
	if err != nil {
		return data{}, err
	}
	err = loadFromToml(fs, "7_checkouts.toml", "checkout", &d.uncashedoutCheckouts)
	if err != nil {
		return data{}, err
	}
	return d, err
}

func loadFromToml[T any](fs fsR, filename, key string, p *T) error {
	dataFile, err := fs.Open(filename)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}
	d := map[string]T{key: *new(T)}
	err = toml.Decode(dataFile, &d)
	if err != nil {
		return err
	}
	*p = d[key]
	return nil
}

func (f *filesystem) SavePersistently() error {
	err1 := f.saveFile("0_user.toml", map[string]any{"user": f.data.users})
	err2 := f.saveFile("1_recipe_groups.toml", map[string]any{"recipe_group": f.data.recipeGroups})
	err3 := f.saveFile("2_recipe_types.toml", map[string]any{"recipe_type": f.data.recipeTypes})
	err4 := f.saveFile("3_glass_types.toml", map[string]any{"glass_type": f.data.glassTypes})
	err5 := f.saveFile("4_menu_items.toml", map[string]any{"menu_items": f.data.menuItems})
	err6 := f.saveFile("5_stocks.toml", map[string]any{"stock": f.data.stocks})
	err7 := f.saveFile("6_orders.toml", map[string]any{"order": f.data.uncheckedOrders})
	err8 := f.saveFile("7_checkouts.toml", map[string]any{"checkout": f.data.uncashedoutCheckouts})
	return errors.Join(err1, err2, err3, err4, err5, err6, err7, err8)
}

func (f *filesystem) saveFile(filename string, data any) error {
	dataFile, err := f.write.Create(filename)
	if err != nil {
		return err
	}
	return toml.Encode(dataFile, data, toml.WithIndent(""))
}
