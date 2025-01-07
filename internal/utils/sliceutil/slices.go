package sliceutil

import "slices"

func Filter[T any](slice []T, match func(T) bool) []T {
	newSlice := make([]T, 0, len(slice))
	for _, value := range slice {
		if match(value) {
			newSlice = append(newSlice, value)
		}
	}
	return newSlice
}

func FilterNonNilPointerValues[T any](slice []*T) []T {
	return Map(
		Filter(slice, func(value *T) bool {
			return value != nil
		}),
		func(value *T) T {
			return *value
		},
	)
}

func Map[T1, T2 any](slice []T1, yield func(T1) T2) []T2 {
	newSlice := make([]T2, 0, len(slice))
	for _, value := range slice {
		newSlice = append(newSlice, yield(value))
	}
	return newSlice
}

func MapE[T1, T2 any](slice []T1, yield func(T1) (T2, error)) ([]T2, error) {
	newSlice := make([]T2, 0, len(slice))
	for _, value := range slice {
		newValue, err := yield(value)
		if err != nil {
			return nil, err
		}
		newSlice = append(newSlice, newValue)
	}
	return newSlice, nil
}

func Some[T any](slice []T, match func(T) bool) bool {
	for _, value := range slice {
		if match(value) {
			return true
		}
	}
	return false
}

func FindOne[T any](slice []T, match func(T) bool) *T {
	for _, value := range slice {
		if match(value) {
			return &value
		}
	}
	return nil
}

func Compact[S ~[]E, E comparable](s S) S {
	return slices.Compact(s)
}

func Reduce[T1, T2 any](slice []T1, initial T2, reduce func(T2, T1) T2) T2 {
	result := initial
	for _, value := range slice {
		result = reduce(result, value)
	}
	return result
}
