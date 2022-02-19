package lists

// Iterator defines the general list iterator interface.
type Iterator[T any] interface {
	HasNext() bool
	Next() T
	Set(T)
	Valid() bool
}

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
	require(it.Valid(), "iterator must be valid")
	require(it.HasNext(), "iterator must have next")

	it.node = it.node.next
	return it.node.value
}

func (it *FrwIter[T]) Set(val T) {
	require(it.Valid(), "iterator must be valid")
	require(it.node != it.lst.head && it.node != it.lst.tail, "iterator must have a value")

	it.node.value = val
}

func (it *FrwIter[T]) Valid() bool {
	return it.node != nil
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
	require(it.Valid(), "iterator must be valid")
	require(it.HasNext(), "iterator must have next")

	it.node = it.node.next
	return it.node.value
}

func (it *RevIter[T]) Set(val T) {
	require(it.Valid(), "iterator must be valid")
	require(it.node != it.lst.head && it.node != it.lst.tail,
		"iterator must have a value")

	it.node.value = val
}

func (it *RevIter[T]) Valid() bool {
	return it.node != nil
}
