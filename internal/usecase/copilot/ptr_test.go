package copilot

func ptr[T any](v T) *T {
	return &v
}
