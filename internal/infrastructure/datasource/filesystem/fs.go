package filesystem

import (
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	"hackbar-copilot/internal/usecase/recipes"
	usecaseutils "hackbar-copilot/internal/usecase/utils"
	"hackbar-copilot/internal/utils/sliceutil"
)

type Filesystem interface {
	recipes.Repository
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
//	├── 1_recipe_groups.toml
//	├── 2_recipe_types.toml
//	└── 3_glass_types.toml
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
	recipeGroups []recipes.RecipeGroup
	recipeTypes  map[string]model.RecipeType
	glassTypes   map[string]model.GlassType
}

func (data data) isEmpty() bool {
	return len(data.recipeGroups) == 0 && len(data.recipeTypes) == 0 && len(data.glassTypes) == 0
}

// Find implements recipes.Repository.
func (f *filesystem) Find() ([]recipes.RecipeGroup, error) {
	return f.data.recipeGroups, nil
}

// FindOne implements recipes.Repository.
func (f *filesystem) FindOne(name string) (recipes.RecipeGroup, error) {
	recipeGroup := sliceutil.FindOne(f.data.recipeGroups, func(rg recipes.RecipeGroup) bool {
		return rg.Name == name
	})
	if recipeGroup == nil {
		return recipes.RecipeGroup{}, usecaseutils.ErrNotFound
	}
	return *recipeGroup, nil
}

// Save implements recipes.Repository.
func (f *filesystem) Save(new recipes.RecipeGroup) error {
	for i, recipeGroup := range f.data.recipeGroups {
		if recipeGroup.Name == new.Name {
			f.data.recipeGroups[i] = new
			return nil
		}
	}
	f.data.recipeGroups = append(f.data.recipeGroups, new)
	return nil
}

// FindRecipeType implements recipes.Repository.
func (f *filesystem) FindRecipeType() (map[string]model.RecipeType, error) {
	return f.data.recipeTypes, nil
}

// FindGlassType implements recipes.Repository.
func (f *filesystem) FindGlassType() (map[string]model.GlassType, error) {
	return f.data.glassTypes, nil
}

// SaveRecipeType implements recipes.Repository.
func (f *filesystem) SaveRecipeType(new model.RecipeType) error {
	f.data.recipeTypes[new.Name] = new
	return nil
}

// SaveGlassType implements recipes.Repository.
func (f *filesystem) SaveGlassType(new model.GlassType) error {
	f.data.glassTypes[new.Name] = new
	return nil
}
