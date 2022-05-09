package gomaps

import (
	"reflect"

	"github.com/shayanh/gcl"
)

// Iterator is an iterator for built-in go maps.
type Iterator[K comparable, V any] struct {
	impl    *reflect.MapIter
	hasNext bool
}

func (it *Iterator[K, V]) HasNext() bool {
	return it.hasNext
}

func (it *Iterator[K, V]) Next() gcl.MapElem[K, V] {
	key := it.impl.Key().Interface().(K)
	value := it.impl.Value().Interface().(V)
	it.hasNext = it.impl.Next()
	return gcl.MapElem[K, V]{
		Key:   key,
		Value: value,
	}
}
