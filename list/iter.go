package list

type Iter[T any] interface {
	Next()
	Prev()
	Done() bool
	Value() T
	SetValue(T)
}

// FrwIter is a forward iterator.
type FrwIter[T any] struct {
	node *node[T]
	lst  *List[T]
}

func (it *FrwIter[T]) Next() {
	it.node = it.node.next
}

func (it *FrwIter[T]) Prev() {
	it.node = it.node.prev
}

func (it *FrwIter[T]) Done() bool {
	return it.node == it.lst.tail || it.node == it.lst.head
}

func (it *FrwIter[T]) Value() T {
	return it.node.value
}

func (it *FrwIter[T]) SetValue(val T) {
	it.node.value = val
}

func (it *FrwIter[T]) clone() *FrwIter[T] {
	return &FrwIter[T]{
		node: it.node,
		lst:  it.lst,
	}
}

// RIter is a reverse iterator.
type RevIter[T any] struct {
	node *node[T]
	lst  *List[T]
}

func (it *RevIter[T]) Next() {
	it.node = it.node.prev
}

func (it *RevIter[T]) Prev() {
	it.node = it.node.next
}

func (it *RevIter[T]) Done() bool {
	return it.node == it.lst.head || it.node == it.lst.tail
}

func (it *RevIter[T]) Value() T {
	return it.node.value
}

func (it *RevIter[T]) SetValue(val T) {
	it.node.value = val
}

func (it *RevIter[T]) clone() *RevIter[T] {
	return &RevIter[T]{
		node: it.node,
		lst:  it.lst,
	}
}
