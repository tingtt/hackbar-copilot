package filesystem

import (
	"errors"
	"hackbar-copilot/internal/infrastructure/datasource/filesystem/toml"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	"hackbar-copilot/internal/usecase/recipes"
	usecaseutils "hackbar-copilot/internal/usecase/utils"
	"hackbar-copilot/internal/utils/sliceutil"
	"io/fs"
	"os"
	"path"
)

// Example usage:
//
//	fs := os.DirFS("/path/to/data-dir")
//	repository := filesystem.NewRepository(fs)
//
// The provided fs.FS needs to follow the structure:
//
//	```
//	.
//	├── manuals
//	│   ├── [uuid]
//	│   └── [uuid]
//	├── orders
//	│   ├── [uuid]
//	│   └── [uuid]
//	├── recipes
//	│   ├── [uuid]
//	│   └── [uuid]
//	```
func NewRepository(baseDir string) (*Filesystem, error) {
	fsR := os.DirFS(baseDir)
	fsW := newFSW(baseDir)
	data, err := loadData(fsR)
	if err != nil {
		return nil, err
	}
	return &Filesystem{read: fsR, write: fsW, data: data}, nil
}

// Checking implements
var _ recipes.Repository = new(Filesystem)

type Filesystem struct {
	read  fs.FS
	write fsW
	data  data
}

type data struct {
	recipeGroups []recipes.RecipeGroup
	recipeTypes  map[string]model.RecipeType
	glassTypes   map[string]model.GlassType
}

func loadData(fs fs.FS) (d data, err error) {
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

func loadRecipeGroups(fs fs.FS) ([]recipes.RecipeGroup, error) {
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

func loadRecipeTypes(fs fs.FS) (map[string]model.RecipeType, error) {
	dataFile, err := fs.Open("2_recipe_types.toml")
	if errors.Is(err, os.ErrNotExist) {
		return make(map[string]model.RecipeType), nil
	}
	if err != nil {
		return make(map[string]model.RecipeType), err
	}
	d := map[string]map[string]model.RecipeType{"recipe_type": {}}
	err = toml.Decode(dataFile, &d)
	if err != nil {
		return make(map[string]model.RecipeType), err
	}
	return d["recipe_type"], nil
}

func loadGlassTypes(fs fs.FS) (map[string]model.GlassType, error) {
	dataFile, err := fs.Open("3_glass_types.toml")
	if errors.Is(err, os.ErrNotExist) {
		return make(map[string]model.GlassType), nil
	}
	if err != nil {
		return make(map[string]model.GlassType), err
	}
	d := map[string]map[string]model.GlassType{"glass_type": {}}
	err = toml.Decode(dataFile, &d)
	if err != nil {
		return make(map[string]model.GlassType), err
	}
	return d["glass_type"], nil
}

func (f *Filesystem) SavePersistently() error {
	{
		dataFile, err := f.write.Create("1_recipe_groups.toml")
		if err != nil {
			return err
		}
		err = toml.Encode(dataFile, map[string]interface{}{"recipe_group": f.data.recipeGroups}, toml.WithIndent(""))
		if err != nil {
			return err
		}
	}
	{
		dataFile, err := f.write.Create("2_recipe_types.toml")
		if err != nil {
			return err
		}
		err = toml.Encode(dataFile, map[string]interface{}{"recipe_type": f.data.recipeTypes}, toml.WithIndent(""))
		if err != nil {
			return err
		}
	}
	{
		dataFile, err := f.write.Create("3_glass_types.toml")
		if err != nil {
			return err
		}
		err = toml.Encode(dataFile, map[string]interface{}{"glass_type": f.data.glassTypes}, toml.WithIndent(""))
		if err != nil {
			return err
		}
	}
	return nil
}

// Find implements recipes.Repository.
func (f *Filesystem) Find() ([]recipes.RecipeGroup, error) {
	return f.data.recipeGroups, nil
}

// FindOne implements recipes.Repository.
func (f *Filesystem) FindOne(name string) (recipes.RecipeGroup, error) {
	recipeGroup := sliceutil.Find(f.data.recipeGroups, func(rg recipes.RecipeGroup) bool {
		return rg.Name == name
	})
	if recipeGroup == nil {
		return recipes.RecipeGroup{}, usecaseutils.ErrNotFound
	}
	return *recipeGroup, nil
}

// Save implements recipes.Repository.
func (f *Filesystem) Save(new recipes.RecipeGroup) error {
	for i, recipeGroup := range f.data.recipeGroups {
		if recipeGroup.Name != new.Name {
			continue
		}
		f.data.recipeGroups[i] = new
		return nil
	}
	f.data.recipeGroups = append(f.data.recipeGroups, new)
	return nil
}

// FindRecipeType implements recipes.Repository.
func (f *Filesystem) FindRecipeType() (map[string]model.RecipeType, error) {
	return f.data.recipeTypes, nil
}

// FindGlassType implements recipes.Repository.
func (f *Filesystem) FindGlassType() (map[string]model.GlassType, error) {
	return f.data.glassTypes, nil
}

// SaveRecipeType implements recipes.Repository.
func (f *Filesystem) SaveRecipeType(new model.RecipeType) error {
	f.data.recipeTypes[new.Name] = new
	return nil
}

// SaveGlassType implements recipes.Repository.
func (f *Filesystem) SaveGlassType(new model.GlassType) error {
	f.data.glassTypes[new.Name] = new
	return nil
}

type fsW interface {
	Create(name string) (*os.File, error)
}

func newFSW(baseDir string) fsW {
	return fsWritable{baseDir}
}

type fsWritable struct {
	baseDir string
}

func (f fsWritable) Create(name string) (*os.File, error) {
	return os.Create(path.Join(f.baseDir, name))
}
