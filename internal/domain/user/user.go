package user

import (
	"errors"
	"iter"
)

type User struct {
	Email         Email
	Name          string
	NameConfirmed bool
}

type Email string

type saver interface {
	Save(d User) error
}

type Lister interface {
	All() iter.Seq2[User, error]
}

var ErrNotFound = errors.New("not found")

type Getter interface {
	Get(email Email) (User, error)
}

type SaveListGetter interface {
	saver
	Lister
	Getter
}

type Repository SaveListGetter

func NewSaveListGetter(r Repository) SaveListGetter {
	return &saveLister{r}
}

type saveLister struct {
	Repository
}
