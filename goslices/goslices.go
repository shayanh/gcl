// Package goslices provides iterators and operations for built-in go slices.
package goslices

import (
	"github.com/shayanh/gcl/internal"
	"github.com/shayanh/gcl/iters"
)

// Iter returns an forward iterator to the beginning. Initially, the returned
// iterator is located at one step before the first element (one-before-first).
func Iter[S ~[]T, T any](s S) *FrwIter[T] {
	return &FrwIter[T]{
		slice: s,
		index: -1,
	}
}

// IterMut returns a forward iterator to the beginning with mutable pointers.
// Initially, the returned iterator is located at one step before the first
// element (one-before-first).
func IterMut[S ~[]T, T any](s S) *FrwIterMut[T] {
	return &FrwIterMut[T]{
		slice: s,
		index: -1,
	}
}

// RIter returns a reverse iterator going from the end to the beginning.
// Initially, the returned iterator is located at one step past the last element
// (one-past-last).
func RIter[S ~[]T, T any](s S) *RevIter[T] {
	return &RevIter[T]{
		slice: s,
		index: len(s),
	}
}

// RIterMut returns a reverse iterator going from the end to the beginning with
// mutable pointers. Initially, the returned iterator is located at one step
// past the last element (one-past-last).
func RIterMut[S ~[]T, T any](s S) *RevIterMut[T] {
	return &RevIterMut[T]{
		slice: s,
		index: len(s),
	}
}

// FromIter builds a new slices from the given iterator.
func FromIter[T any](it iters.Iterator[T]) (res []T) {
	for it.HasNext() {
		res = append(res, it.Next())
	}
	return
}

// PushFront appends the given elements to the beginning of slice s
// and returns the resulting slice.
// This function is O(len(s) + len(elems)).
func PushFront[S ~[]T, T any](s S, elems ...T) S {
	return append(elems, s...)
}

// PopBack deletes the last element in a slice and returns the resulting
// slice. PopBack requires the given slice to be non-empty.
// This function is O(1).
func PopBack[S ~[]T, T any](s S) S {
	return s[0 : len(s)-1]
}

// PopFront deletes the first element in a slice and returns the resulting
// slices. PopFront requires the slice to be non-empty, otherwise it panics.
// This function is O(1).
func PopFront[S ~[]T, T any](s S) S {
	return s[1:]
}

// Front returns the first element in a slice. It panics if the given slice is
// empty. This function is O(1).
func Front[S ~[]T, T any](s S) T {
	return s[0]
}

// Front returns the last element in a slice. It panics if the given slice is
// empty. This function is O(1).
func Back[S ~[]T, T any](s S) T {
	return s[len(s)-1]
}

// Reverse reverses the elements of a slice and returns the resulting slice.
// This function is O(n), where n is length of the given slice.
func Reverse[S ~[]T, T any](s S) {
	internal.Reverse[T](IterMut(s), RIterMut(s), len(s))
}
