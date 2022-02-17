package iters

import (
	"constraints"

	"github.com/shayanh/gogl"
)

func ForEach[T any](it Iterator[T], fn func(T)) {
	for it.HasNext() {
		fn(it.Next())
	}
}

type mapIter[T any, V any] struct {
	wrappedIt Iterator[T]
	fn func(T) V
}

func (it *mapIter[T, V]) HasNext() bool {
	return it.wrappedIt.HasNext()
}

func (it *mapIter[T, V]) Next() V {
	return it.fn(it.wrappedIt.Next())
}

func Map[T any, V any](it Iterator[T], fn func(T) V) Iterator[V] {
	return &mapIter[T, V]{wrappedIt: it, fn: fn}
}

func Reduce[T any](it Iterator[T], fn func(T, T) T) (acc T) {
	if !it.HasNext() {
		return
	}
	acc = it.Next()
	for it.HasNext() {
		acc = fn(acc, it.Next())
	}
	return
}

func Fold[T any, V any](it Iterator[T], fn func(V, T) V, init V) (acc V) {
	acc = init
	for it.HasNext() {
		acc = fn(acc, it.Next())
	}
	return
}

type nextState int

const (
	unknown nextState = iota 
	hasNext
	noNext
)

type filterIter[T any] struct {
	wrappedIt Iterator[T]
	pred func(T) bool
	state nextState	
	next T
}

func (it *filterIter[T]) findNext() {
	for it.wrappedIt.HasNext() {
		v := it.wrappedIt.Next()
		if it.pred(v) {
			it.state = hasNext
			it.next = v
			return
		}
	}
	it.state = noNext
}

func (it *filterIter[T]) HasNext() bool {
	if it.state == unknown {
		it.findNext()
	}
	if it.state == hasNext {
		return true
	}
	return false
}

func (it *filterIter[T]) Next() T {
	if it.state == unknown {
		it.findNext()
	}
	if it.state == noNext {
		panic("iterator does not have next")
	}
	it.state = unknown
	return it.next
}

func Filter[T any](it Iterator[T], pred func(T) bool) Iterator[T] {
	return &filterIter[T]{wrappedIt: it, pred: pred}
}

func Max[T constraints.Ordered](it Iterator[T]) (max T) {
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

func MaxFunc[T any](it Iterator[T], less gogl.LessFn[T]) (max T) {
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
