package goslices

// FrwIter is a slice forward iterator.
type FrwIter[T any] struct {
	slice []T
	index int
}

func (it *FrwIter[T]) HasNext() bool {
	return it.index+1 < len(it.slice)
}

func (it *FrwIter[T]) Next() T {
	it.index += 1
	return it.slice[it.index]
}

// FrwIterMut is a mutable forward iterator for slices. It allows mutations by
// returning pointers to the slice elements. FrwIterMut is an iterator over
// pointers of type T. In other words, FrwIterMut[T] implements
// iters.Iterator[*T].
type FrwIterMut[T any] struct {
	slice []T
	index int
}

func (it *FrwIterMut[T]) HasNext() bool {
	return it.index+1 < len(it.slice)
}

func (it *FrwIterMut[T]) Next() *T {
	it.index += 1
	return &it.slice[it.index]
}

// RevIter is a slice reverse iterator.
type RevIter[T any] struct {
	slice []T
	index int
}

func (it *RevIter[T]) HasNext() bool {
	return it.index > 0
}

func (it *RevIter[T]) Next() T {
	it.index -= 1
	return it.slice[it.index]
}

// RevIterMut is a mutable reverse iterator for slices. It allows mutations by
// returning pointers to the slice elements. RevIterMut is an iterator over
// pointers of type T. In other words, RevIterMut[T] implements
// iters.Iterator[*T].
type RevIterMut[T any] struct {
	slice []T
	index int
}

func (it *RevIterMut[T]) HasNext() bool {
	return it.index > 0
}

func (it *RevIterMut[T]) Next() *T {
	it.index -= 1
	return &it.slice[it.index]
}
