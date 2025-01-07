package sort

type Yield[T any] func(new, curr T) (isLeft bool)
type YieldMaker[T any] func(fallbacks ...YieldMaker[T]) Yield[T]

func Desc[T any](yield Yield[T]) Yield[T] {
	return func(new, curr T) (isLeft bool) {
		return !yield(new, curr)
	}
}
