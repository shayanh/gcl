package internal

import (
	"github.com/shayanh/gogl"
	"github.com/shayanh/gogl/iters"
)

func ForEach[T any](it iters.Iter[T], fn func(T)) {
	for ; !it.Done(); it.Next() {
		fn(it.Value())
	}
}

func Reverse[T any](fIt iters.MutIter[T], rIt iters.MutIter[T], length int) {
	fIdx, rIdx := 0, length-1
	for fIdx < rIdx {
		fVal, rVal := fIt.Value(), rIt.Value()
		fIt.SetValue(rVal)
		rIt.SetValue(fVal)

		fIdx += 1
		rIdx -= 1
		fIt.Next()
		rIt.Next()

		if fIt.Done() || rIt.Done() {
			panic("bad iterator")
		}
	}
}

func MaxFunc[T any](it iters.Iter[T], less gogl.LessFn[T]) (max T) {
	if it.Done() {
		return
	}

	max = it.Value()
	for ; !it.Done(); it.Next() {
		if less(max, it.Value()) {
			max = it.Value()
		}
	}
	return
}

func MinFunc[T any](it iters.Iter[T], less gogl.LessFn[T]) (min T) {
	if it.Done() {
		return
	}

	min = it.Value()
	for ; !it.Done(); it.Next() {
		if less(it.Value(), min) {
			min = it.Value()
		}
	}
	return
}
