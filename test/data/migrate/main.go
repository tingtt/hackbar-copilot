package main

import (
	"fmt"
	"hackbar-copilot/internal/domain/menu/menutest"
	"hackbar-copilot/internal/domain/order/ordertest"
	"hackbar-copilot/internal/domain/recipe/recipetest"
	"hackbar-copilot/internal/domain/stock/stocktest"
	"hackbar-copilot/internal/domain/user/usertest"
	"hackbar-copilot/internal/infrastructure/datasource/filesystem"
	"log/slog"
	"os"
	"path"

	"github.com/spf13/pflag"
)

func main() {
	dir := pflag.String("dir", "./.data/", "directory to migrate data")
	clear := pflag.Bool("clear", false, "clear directory before migrating")
	pflag.Parse()

	err := run(*dir, *clear)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func run(dir string, clear bool) error {
	if clear {
		slog.Info("clearing exists data...")
		err := clearDir(dir)
		if err != nil {
			return err
		}
		slog.Info("done (clearing)")
	}

	slog.Info("migrating data...")
	err := migrate(dir)
	if err != nil {
		return err
	}
	slog.Info("done (migrating)")
	return nil
}

func clearDir(dataDir string) error {
	entries, err := os.ReadDir(dataDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if path.Ext(entry.Name()) == ".toml" {
			err = os.Remove(path.Join(dataDir, entry.Name()))
			if err != nil {
				return fmt.Errorf("failed to remove file %s: %w", entry.Name(), err)
			}
		}
	}

	return nil
}

func migrate(dataDir string) error {
	fs, err := filesystem.NewRepository(dataDir)
	if err != nil {
		return err
	}

	err = registerData(fs)
	if err != nil {
		return err
	}

	return fs.SavePersistently()
}

func registerData(fs filesystem.Filesystem) error {
	for user := range usertest.ExampleUsersIter {
		err := fs.OrderGateway().User().Save(user)
		if err != nil {
			return err
		}
	}
	for recipeGroup := range recipetest.ExampleRecipeGroupsIter {
		err := fs.CopilotGateway().Recipe().Save(recipeGroup)
		if err != nil {
			return err
		}
	}
	for recipeType := range recipetest.ExampleRecipeTypesIter {
		err := fs.CopilotGateway().Recipe().SaveRecipeType(recipeType)
		if err != nil {
			return err
		}
	}
	for glassType := range recipetest.ExampleGlassTypesIter {
		err := fs.CopilotGateway().Recipe().SaveGlassType(glassType)
		if err != nil {
			return err
		}
	}

	for menuGroup := range menutest.ExampleItemsIter {
		err := fs.CopilotGateway().Menu().Save(menuGroup)
		if err != nil {
			return err
		}
	}

	inStockMaterialNames, outOfStockMaterialNames := stocktest.ExampleMaterialNames()
	err := fs.CopilotGateway().Stock().Save(inStockMaterialNames, outOfStockMaterialNames)
	if err != nil {
		return err
	}

	for order := range ordertest.ExampleOrdersIter {
		err := fs.OrderGateway().Order().Save(order)
		if err != nil {
			return err
		}
	}

	return nil
}
