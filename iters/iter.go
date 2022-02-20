// Package iters defines the general iterator interface and provides operations
// on the given iterators.
package iters

// Iterator defines the general iterator interface. Iterator makes it
// possible to iterate over collections.
type Iterator[T any] interface {
	// HasNext tests whether the iterator can advance.
	HasNext() bool

	// Next advances the iterator and returns the next value in iteration.
	Next() T
}

// MutIterator defines the general interface for an iterator that
// its values can be mutated.
type MutIterator[T any] interface {
	Iterator[T]

	// Set replaces the last value returned by Next() with the given value.
	Set(T)
}
