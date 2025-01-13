package menu

import "iter"

type SaveLister interface {
	saver
	Lister
}

type saver interface {
	Save(g Group) error
}

type Lister interface {
	All() iter.Seq2[Group, error]
}

type Repository SaveLister

func NewLister(r Repository) Lister {
	return &saveLister{r}
}

func NewSaveLister(r Repository) SaveLister {
	return &saveLister{r}
}

type saveLister struct {
	Repository
}
