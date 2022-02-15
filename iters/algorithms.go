package iters

import (
	"constraints"

	"github.com/shayanh/gogl"
)

func ForEach[T any](it Iter[T], fn func(T)) {
	for it.HasNext() {
		fn(it.Next())
	}
}

func Map[T any, V any](it Iter[T], fn func(T) V) []V {
	var res []V
	for it.HasNext() {
		res = append(res, fn(it.Next()))
	}
	return res
}

func Reduce[T any](it Iter[T], fn func(T, T) T) (acc T) {
	if !it.HasNext() {
		return
	}
	acc = it.Next()
	for it.HasNext() {
		acc = fn(acc, it.Next())
	}
	return
}

func Fold[T any, V any](it Iter[T], fn func(V, T) V, init V) (acc V) {
	acc = init
	for it.HasNext() {
		acc = fn(acc, it.Next())
	}
	return
}

func Filter[T any](it Iter[T], pred func(T) bool) []T {
	var res []T
	for it.HasNext() {
		v := it.Next()
		if pred(v) {
			res = append(res, v)
		}
	}
	return res
}

func Max[T constraints.Ordered](it Iter[T]) (max T) {
	if !it.HasNext() {
		return
	}

	max = it.Next()
	for it.HasNext() {
		v := it.Next()
		if max < v {
			max = v
		}
	}
	return
}

func MaxFunc[T any](it Iter[T], less gogl.LessFn[T]) (max T) {
	if !it.HasNext() {
		return
	}

	max = it.Next()
	for it.HasNext() {
		v := it.Next()
		if less(max, v) {
			max = v
		}
	}
	return
}
