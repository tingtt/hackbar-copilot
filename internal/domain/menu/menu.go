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
	Save(g Group) error
}

var ErrNotFound = errors.New("menu not found")

type Finder interface {
	Find(groupName, itemName string) (Item, error)
}

type Lister interface {
	All() iter.Seq2[Group, error]
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
