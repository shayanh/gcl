package goslices

type Iterator[T any] interface {
	HasNext() bool
	Next() T
	Set(T)
}

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

func (it *FrwIter[T]) Set(val T) {
	it.slice[it.index] = val
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

func (it *RevIter[T]) Set(val T) {
	it.slice[it.index] = val
}
