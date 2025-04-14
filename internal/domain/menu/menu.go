package menu

import (
	"errors"
	"iter"
)

type SaveFindListRemover interface {
	saver
	Finder
	Lister
	Remover
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

type Remover interface {
	Remove(itemName string) error
}

type Repository SaveFindListRemover

func NewFindLister(r Repository) FindLister {
	return &saveFindLister{r}
}

func NewSaveLister(r Repository) SaveFindListRemover {
	return &saveFindLister{r}
}

type saveFindLister struct {
	Repository
}
