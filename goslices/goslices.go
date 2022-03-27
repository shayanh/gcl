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

func Reverse[S ~[]T, T any](s S) {
	internal.Reverse[T](Iter(s), RIter(s), len(s))
}
