package gomaps

func Iter[M ~map[K]V, K comparable, V any](m M) *Iterator[K, V] {
	return &Iterator[K, V]{
		m:     m,
		index: -1,
	}
}
