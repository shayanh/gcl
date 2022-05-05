package lists

func require(check bool, failMsg string) {
	if !check {
		panic(failMsg)
	}
}

// FrwIter is a forward iterator.
type FrwIter[T any] struct {
	node *node[T]
	lst  *List[T]
}

func (it *FrwIter[T]) HasNext() bool {
	if it.node == nil {
		return false
	}
	return it.node.next != it.lst.tail
}

func (it *FrwIter[T]) Next() T {
	require(it.HasNext(), "iterator must have next")

	it.node = it.node.next
	return it.node.value
}

type FrwIterMut[T any] struct {
	node *node[T]
	lst  *List[T]
}

func (it *FrwIterMut[T]) HasNext() bool {
	if it.node == nil {
		return false
	}
	return it.node.next != it.lst.tail
}

func (it *FrwIterMut[T]) Next() *T {
	require(it.HasNext(), "iterator must have next")

	it.node = it.node.next
	return &it.node.value
}

func (it *FrwIterMut[T]) Insert(elems ...T) {
	require(it.node.next != nil, "bad iterator")
	for i := len(elems) - 1; i >= 0; i-- {
		elem := elems[i]
		node := &node[T]{value: elem}
		it.lst.insertBetween(node, it.node, it.node.next)
	}
}

func (it *FrwIterMut[T]) Delete() {
	require(it.node.prev != nil && it.node.next != nil,
		"bad iterator")
	prev, _ := it.lst.deleteNode(it.node)
	it = &FrwIterMut[T]{
		node: prev,
		lst:  it.lst,
	}
}

// RevIter is a reverse iterator.
type RevIter[T any] struct {
	node *node[T]
	lst  *List[T]
}

func (it *RevIter[T]) HasNext() bool {
	if it.node == nil {
		return false
	}
	return it.node.prev != it.lst.head
}

func (it *RevIter[T]) Next() T {
	require(it.HasNext(), "iterator must have next")

	it.node = it.node.prev
	return it.node.value
}

type RevIterMut[T any] struct {
	node *node[T]
	lst  *List[T]
}

func (it *RevIterMut[T]) HasNext() bool {
	if it.node == nil {
		return false
	}
	return it.node.prev != it.lst.head
}

func (it *RevIterMut[T]) Next() *T {
	require(it.HasNext(), "iterator must have next")

	it.node = it.node.prev
	return &it.node.value
}

func (it *RevIterMut[T]) Insert(elems ...T) {
	require(it.node.prev != nil, "bad iterator")
	for _, elem := range elems {
		node := &node[T]{value: elem}
		it.lst.insertBetween(node, it.node.prev, it.node)
	}
}

func (it *RevIterMut[T]) Delete() {
	require(it.node.prev != nil && it.node.next != nil,
		"bad iterator")
	_, next := it.lst.deleteNode(it.node)
	it = &RevIterMut[T]{
		node: next,
		lst:  it.lst,
	}
}
