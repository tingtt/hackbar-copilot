package usertest

import (
	"hackbar-copilot/internal/domain/user"
	"iter"
)

func IterWithNilError(items []user.User) iter.Seq2[user.User, error] {
	return func(yield func(user.User, error) bool) {
		for _, item := range items {
			if !yield(item, nil) {
				break
			}
		}
	}
}
