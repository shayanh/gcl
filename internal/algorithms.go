package internal

import (
	"github.com/shayanh/gogl/iters"
)

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
