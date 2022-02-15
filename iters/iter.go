package iters

type Iterator[T any] interface {
	// HasNext tests whether the iterator can advance.
	HasNext() bool

	// Next advances the iterator and returns the next value in iteration.
	Next() T
}

type MutIterator[T any] interface {
	Iterator[T]

	// Set replaces the last value returned by Next with the given value.
	Set(T)
}
