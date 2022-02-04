package list

// Iter is a forward iterator.
type Iter[T any] struct {
	node *node[T]
	lst  *List[T]
}

func (it *Iter[T]) Next() {
	it.node = it.node.next
}

func (it *Iter[T]) Prev() {
	it.node = it.node.prev
}

func (it *Iter[T]) Done() bool {
	return it.node == it.lst.back
}

func (it *Iter[T]) Value() T {
	return it.node.value
}

func (it *Iter[T]) SetValue(val T) {
	it.node.value = val
}

// RIter is a reverse iterator.
type RIter[T any] struct {
	node *node[T]
	lst  *List[T]
}

func (it *RIter[T]) Next() {
	it.node = it.node.prev
}

func (it *RIter[T]) Prev() {
	it.node = it.node.next
}

func (it *RIter[T]) Done() bool {
	return it.node == it.lst.front
}

func (it *RIter[T]) Value() T {
	return it.node.value
}

func (it *RIter[T]) SetValue(val T) {
	it.node.value = val
}
