package internal

import (
	"github.com/shayanh/gogl/iters"
)

func Reverse[T any](fIt iters.MutIter[T], rIt iters.MutIter[T], length int) {
	fIdx, rIdx := 0, length-1
	for fIdx < rIdx {
		if !fIt.HasNext() || !rIt.HasNext() {
			panic("bad iterator")
		}

		fVal, rVal := fIt.Next(), rIt.Next()
		fIt.Set(rVal)
		rIt.Set(fVal)

		fIdx += 1
		rIdx -= 1
	}
}
