package recipe

import (
	"iter"
)

type SaveLister interface {
	Saver
	Lister
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

type Repository SaveLister

func NewSaveLister(r Repository) SaveLister {
	return &saverLister{r}
}

var _ SaveLister = new(saverLister)

type saverLister struct {
	Repository
}
