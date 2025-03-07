package menu

import (
	"errors"
	"iter"
)

type SaveFindLister interface {
	saver
	Finder
	Lister
}

type FindLister interface {
	Finder
	Lister
}

type saver interface {
	Save(g Item) error
}

var ErrNotFound = errors.New("menu not found")

type Finder interface {
	Find(itemName, optionName string) (ItemOption, error)
}

type Lister interface {
	All() iter.Seq2[Item, error]
}

type Repository SaveFindLister

func NewFindLister(r Repository) FindLister {
	return &saveFindLister{r}
}

func NewSaveLister(r Repository) SaveFindLister {
	return &saveFindLister{r}
}

type saveFindLister struct {
	Repository
}
