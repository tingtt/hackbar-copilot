package recipetest

import "iter"

func ptr[T any](v T) *T {
	return &v
}

func iterWithNilError[T any](items []T) iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		for _, item := range items {
			if !yield(item, nil) {
				break
			}
		}
	}
}
