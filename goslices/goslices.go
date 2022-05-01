package goslices

import (
	"github.com/shayanh/gcl/internal"
	"github.com/shayanh/gcl/iters"
)

func Iter[S ~[]T, T any](s S) Iterator[T] {
	return &FrwIter[T]{
		slice: s,
		index: -1,
	}
}

func RIter[S ~[]T, T any](s S) Iterator[T] {
	return &RevIter[T]{
		slice: s,
		index: len(s),
	}
}

func FromIter[T any](it iters.Iterator[T]) (res []T) {
	for it.HasNext() {
		res = append(res, it.Next())
	}
	return
}

func PushFront[S ~[]T, T any](s S, elems ...T) S {
	return append(elems, s...)
}

func PopBack[S ~[]T, T any](s S) S {
	return s[0 : len(s)-1]
}

func PopFront[S ~[]T, T any](s S) S {
	return s[1:]
}

func Front[S ~[]T, T any](s S) T {
	return s[0]
}

func Back[S ~[]T, T any](s S) T {
	return s[len(s)-1]
}

func Reverse[S ~[]T, T any](s S) {
	internal.Reverse[T](Iter(s), RIter(s), len(s))
}
