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

func (it *FrwIter) HasNext() bool {
	return it.index + 1 < len(it.slice)
}

func (it *FrwIter) Next() T {
	it.index += 1
	return it.slice[index]
}

func (it *FrwIter) Set(val T) {
	it.slice[it.index] = val
}

type RevIter[T any] struct {
	HasNext() bool
	Next() T
	Set(T)
}

func (it *RevIter) HasNext() bool {
	return it.index > 0
}

func (it *RevIter) Next() T {
	it.index -= 1
	return it.slice[index]
}

func (it *RevIter) Set(val T) {
	it.slice[it.index] = val
}
