package internal

import (
	"github.com/shayanh/gogl/iter"
)

func ForEach[T any](it iter.Iter[T], fn func(T)) {
	for ; !it.Done(); it.Next() {
		fn(it.Value())
	}
}

func Reverse[T any](fIt iter.MutIter[T], rIt iter.MutIter[T], length int) {
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
