package iters

func Map[T any, V any](it Iter[T], mapFn func(T) V, insertFn func(V)) {
	for ; !it.Done(); it.Next() {
		insertFn(mapFn(it.Value()))
	}
}
