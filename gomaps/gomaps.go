// Package gomaps provides iterators for built-in go maps.
package gomaps

import (
	"reflect"

	"github.com/shayanh/gcl"
	"github.com/shayanh/gcl/iters"
)

// Iter returns an iterator for the given map.
func Iter[M ~map[K]V, K comparable, V any](m M) *Iterator[K, V] {
	impl := reflect.ValueOf(m).MapRange()
	return &Iterator[K, V]{
		impl:    impl,
		hasNext: impl.Next(),
	}
}

// FromIter builds a new map from an iterator.
func FromIter[K comparable, V any, IT iters.Iterator[gcl.MapElem[K, V]]](it IT) (res map[K]V) {
	res = make(map[K]V)
	for it.HasNext() {
		elem := it.Next()
		res[elem.Key] = elem.Value
	}
	return
}
