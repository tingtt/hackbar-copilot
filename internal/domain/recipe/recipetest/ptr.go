package recipetest

func ptr[T any](v T) *T {
	return &v
}
