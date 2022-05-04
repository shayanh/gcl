package internal

import (
	"github.com/shayanh/gcl/iters"
)

func Reverse[T any](fIt iters.Iterator[*T], rIt iters.Iterator[*T], length int) {
	fIdx, rIdx := 0, length-1
	for fIdx < rIdx {
		if !fIt.HasNext() || !rIt.HasNext() {
			panic("bad iterator")
		}

		fVal, rVal := fIt.Next(), rIt.Next()

		tmp := *fVal
		*fVal = *rVal
		*rVal = tmp

		fIdx += 1
		rIdx -= 1
	}
}
