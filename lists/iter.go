package lists

func require(check bool, failMsg string) {
	if !check {
		panic(failMsg)
	}
}

// FrwIter is a list forward iterator.
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

// FrwIterMut is a mutable forward iterator for lists. It allows mutations by
// returning pointers to the list elements. FrwIterMut is an iterator over
// pointers of type T. In other words, FrwIterMut[T] implements iters.Iterator[*T].
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

// Insert inserts the given values next after the iterator it. This function is
// O(len(elems)). So inserting a single element would be O(1).
func (it *FrwIterMut[T]) Insert(elems ...T) {
	require(it.node.next != nil, "bad iterator")
	for i := len(elems) - 1; i >= 0; i-- {
		elem := elems[i]
		node := &node[T]{value: elem}
		it.lst.insertBetween(node, it.node, it.node.next)
	}
}

// Delete deletes the element that the iterator it is pointing to. Delete
// requires the iterator it to point to an actual element in the list. For
// example, it's not possible to call Delete on the IterMut iterator (in its
// inital state) because this iterator is located at one step before the first
// element and this is not an actual list element. This function is O(1).
func (it *FrwIterMut[T]) Delete() {
	require(it.node.prev != nil && it.node.next != nil, "bad iterator")
	prev, _ := it.lst.deleteNode(it.node)
	it = &FrwIterMut[T]{
		node: prev,
		lst:  it.lst,
	}
}

// RevIter is a list reverse iterator.
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

// RevIterMut is a mutable reverse iterator for lists. It allows mutations by
// returning pointers to the list elements. RevIterMut is an iterator over
// pointers of type T. In other words, RevIterMut[T] implements iters.Iterator[*T].
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

// Insert inserts the given values next after (in a reverse direction) the
// iterator it. This function is O(len(elems)). So inserting a single element
// would be O(1).
func (it *RevIterMut[T]) Insert(elems ...T) {
	require(it.node.prev != nil, "bad iterator")
	for _, elem := range elems {
		node := &node[T]{value: elem}
		it.lst.insertBetween(node, it.node.prev, it.node)
	}
}

// Delete deletes the element that the iterator it is pointing to. Delete
// requires the iterator it to point to an actual element in the list. For
// example, it's not possible to call Delete on the RIterMut iterator (in its
// inital state) because this iterator is located at one step past the last
// element and this is not an actual list element. This function is O(1).
func (it *RevIterMut[T]) Delete() {
	require(it.node.prev != nil && it.node.next != nil,
		"bad iterator")
	_, next := it.lst.deleteNode(it.node)
	it = &RevIterMut[T]{
		node: next,
		lst:  it.lst,
	}
}
