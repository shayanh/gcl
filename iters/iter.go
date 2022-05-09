// Package iters defines the general iterator interface and provides different
// operations on top of them.
package iters

// Iterator defines the general iterator interface. Iterator makes it
// possible to iterate over collections.
type Iterator[T any] interface {
	// HasNext tests whether the iterator can advance.
	HasNext() bool

	// Next advances the iterator and returns the next value in iteration.
	Next() T
}

// Advance advances an iterator n steps. Advance stops at any point
// where the given iterator doesn't have a next element.
func Advance[T any](it Iterator[T], n uint) {
	var i uint
	for i = 0; i < n && it.HasNext(); i++ {
		it.Next()
	}
}
