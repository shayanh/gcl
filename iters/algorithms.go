package iters

func ForEach[T any](cit Iter[T], fn func(T)) {
	it := cit.Clone()
	for ; !it.Done(); it.Next() {
		fn(it.Value())
	}
}

func Map[T any, V any](cit Iter[T], fn func(T) V) []V {
	it := cit.Clone()
	var res []V
	for ; !it.Done(); it.Next() {
		res = append(res, fn(it.Value()))
	}
	return res
}

func Reduce[T any](cit Iter[T], fn func(T, T) T) (acc T) {
	it := cit.Clone()
	if it.Done() {
		return
	}
	acc = it.Value()
	for it.Next(); !it.Done(); it.Next() {
		acc = fn(acc, it.Value())
	}
	return
}

func Fold[T any, V any](cit Iter[T], fn func(V, T) V, init V) (acc V) {
	it := cit.Clone()
	acc = init
	for ; !it.Done(); it.Next() {
		acc = fn(acc, it.Value())
	}
	return
}

func Filter[T any](cit Iter[T], pred func(T) bool) []T {
	it := cit.Clone()
	var res []T
	for ; !it.Done(); it.Next() {
		if pred(it.Value()) {
			res = append(res, it.Value())
		}
	}
	return res
}
