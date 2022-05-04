package goslices

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
