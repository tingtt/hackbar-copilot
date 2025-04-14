package recipe

import (
	"errors"
	"iter"
)

type SaveListRemover interface {
	Saver
	Lister
	Remover
}

type Saver interface {
	Save(rg RecipeGroup) error
	SaveRecipeType(rt RecipeType) error
	SaveGlassType(gt GlassType) error
}

type Lister interface {
	All() iter.Seq2[RecipeGroup, error]
	AllRecipeTypes() iter.Seq2[RecipeType, error]
	AllGlassTypes() iter.Seq2[GlassType, error]
}

var ErrNotFound = errors.New("menu not found")

type Remover interface {
	Remove(recipeGroupName string) error
}

type Repository SaveListRemover

func NewSaveLister(r Repository) SaveListRemover {
	return &saverLister{r}
}

var _ SaveListRemover = new(saverLister)

type saverLister struct {
	Repository
}
