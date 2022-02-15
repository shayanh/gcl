package iters

type Iter[T any] interface {
	// HasNext tests whether the iterator can advance.
	HasNext() bool

	// Next advances the iterator and returns the next value in iteration.
	Next() T
}

type MutIter[T any] interface {
	Iter[T]

	// Set replaces the last value returned by Next with the given value.
	Set(T)
}
