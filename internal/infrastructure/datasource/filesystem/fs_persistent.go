package filesystem

import (
	"errors"
	"hackbar-copilot/internal/infrastructure/datasource/filesystem/toml"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	"hackbar-copilot/internal/usecase/recipes"
	"os"
)

func loadData(fs fsR) (d data, err error) {
	d.recipeGroups, err = loadRecipeGroups(fs)
	if err != nil {
		return data{}, err
	}
	d.recipeTypes, err = loadRecipeTypes(fs)
	if err != nil {
		return data{}, err
	}
	d.glassTypes, err = loadGlassTypes(fs)
	if err != nil {
		return data{}, err
	}
	return d, err
}

func loadRecipeGroups(fs fsR) ([]recipes.RecipeGroup, error) {
	dataFile, err := fs.Open("1_recipe_groups.toml")
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	d := map[string][]recipes.RecipeGroup{"recipe_group": {}}
	err = toml.Decode(dataFile, &d)
	if err != nil {
		return nil, err
	}
	return d["recipe_group"], nil
}

func loadRecipeTypes(fs fsR) (map[string]model.RecipeType, error) {
	dataFile, err := fs.Open("2_recipe_types.toml")
	if errors.Is(err, os.ErrNotExist) {
		return make(map[string]model.RecipeType), nil
	}
	if err != nil {
		return nil, err
	}
	d := map[string]map[string]model.RecipeType{"recipe_type": {}}
	err = toml.Decode(dataFile, &d)
	if err != nil {
		return nil, err
	}
	return d["recipe_type"], nil
}

func loadGlassTypes(fs fsR) (map[string]model.GlassType, error) {
	dataFile, err := fs.Open("3_glass_types.toml")
	if errors.Is(err, os.ErrNotExist) {
		return make(map[string]model.GlassType), nil
	}
	if err != nil {
		return nil, err
	}
	d := map[string]map[string]model.GlassType{"glass_type": {}}
	err = toml.Decode(dataFile, &d)
	if err != nil {
		return nil, err
	}
	return d["glass_type"], nil
}

func (f *filesystem) SavePersistently() error {
	err1 := f.saveFile("1_recipe_groups.toml", map[string]interface{}{"recipe_group": f.data.recipeGroups})
	err2 := f.saveFile("2_recipe_types.toml", map[string]interface{}{"recipe_type": f.data.recipeTypes})
	err3 := f.saveFile("3_glass_types.toml", map[string]interface{}{"glass_type": f.data.glassTypes})
	return errors.Join(err1, err2, err3)
}

func (f *filesystem) saveFile(filename string, data any) error {
	dataFile, err := f.write.Create(filename)
	if err != nil {
		return err
	}
	return toml.Encode(dataFile, data, toml.WithIndent(""))
}
