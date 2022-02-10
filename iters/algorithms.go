package iters

func ForEach[T any](it Iter[T], fn func(T)) {
	for ; !it.Done(); it.Next() {
		fn(it.Value())
	}
}

func Map[T any, V any](it Iter[T], mapFn func(T) V) []V {
	var res []V
	for ; !it.Done(); it.Next() {
		res = append(res, mapFn(it.Value()))
	}
	return res
}
