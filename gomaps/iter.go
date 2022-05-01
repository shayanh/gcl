package gomaps

import (
	"github.com/shayanh/gcl"
	"golang.org/x/exp/maps"
)

type Iterator[K comparable, V any] struct {
	m     map[K]V
	keys  []K
	index int
}

func (it *Iterator[K, V]) HasNext() bool {
	if it.keys == nil {
		return len(it.m) > 0
	}
	return it.index+1 < len(it.keys)
}

func (it *Iterator[K, V]) Next() gcl.MapElem[K, V] {
	if it.keys == nil {
		it.keys = maps.Keys(it.m)
	}
	it.index += 1
	key := it.keys[it.index]
	return gcl.MapElem[K, V]{
		Key:   key,
		Value: it.m[key],
	}
}
