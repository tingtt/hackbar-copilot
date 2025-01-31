package stock

import "iter"

type SaveLister interface {
	saver
	Lister
}

type saver interface {
	Save(inStockMaterials, outOfStockMaterials []string) error
}

type Lister interface {
	All() iter.Seq2[Material, error]
}

type Repository SaveLister

func NewSaveLister(r Repository) SaveLister {
	return &saveLister{r}
}

type saveLister struct {
	Repository
}
