package gomaps

import (
	"github.com/shayanh/gcl"
	"github.com/shayanh/gcl/iters"
)

func Iter[M ~map[K]V, K comparable, V any](m M) *Iterator[K, V] {
	return &Iterator[K, V]{
		m:     m,
		index: -1,
	}
}

func FromIter[K comparable, V any](it iters.Iterator[gcl.MapElem[K, V]]) (res map[K]V) {
	for it.HasNext() {
		elem := it.Next()
		res[elem.Key] = elem.Value
	}
	return
}
