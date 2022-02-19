// Package list provides a doubly linked list.

package lists

import (
	"github.com/shayanh/gogl"
	"github.com/shayanh/gogl/iters"
	"github.com/shayanh/gogl/internal"
)

type node[T any] struct {
	value T
	next  *node[T]
	prev  *node[T]
}

// List is doubly linked list.
type List[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

func NewList[T any](elems ...T) *List[T] {
	head := &node[T]{}
	tail := &node[T]{}

	head.next = tail
	tail.prev = head
	l := &List[T]{
		head: head,
		tail: tail,
	}

	PushBack(l, elems...)
	return l
}

// FromIter builds a new list from the given iterator.
func FromIter[T any](it iters.Iterator[T]) *List[T] {
	l := NewList[T]()
	for it.HasNext() {
		PushBack(l, it.Next())
	}
	return l
}

// Len returns size of the given list.
func Len[T any](l *List[T]) int {
	return l.size
}

// Iter returns a forward iterator to the beginning.
func Iter[T any](l *List[T]) Iterator[T] {
	return &FrwIter[T]{
		node: l.head,
		lst:  l,
	}
}

// RIter returns a reverse iterator to the beginning (in the reverse order).
func RIter[T any](l *List[T]) Iterator[T] {
	return &RevIter[T]{
		node: l.tail,
		lst:  l,
	}
}

func Equal[T comparable](l1, l2 *List[T]) bool {
	if l1.size != l2.size {
		return false
	}
	it1, it2 := Iter(l1), Iter(l2)
	for it1.HasNext() {
		v1 := it1.Next()
		v2 := it2.Next()
		if v1 != v2 {
			return false
		}
	}
	return true
}

func EqualFunc[T any](l1, l2 *List[T], eq gogl.EqualFn[T]) bool {
	if l1.size != l2.size {
		return false
	}
	it1, it2 := Iter(l1), Iter(l2)
	for it1.HasNext() {
		v1 := it1.Next()
		v2 := it2.Next()
		if !eq(v1, v2) {
			return false
		}
	}
	return true
}

func PushBack[T any](l *List[T], elems ...T) {
	Insert(RIter(l), elems...)
}

func PushFront[T any](l *List[T], elems ...T) {
	Insert(Iter(l), elems...)
}

func PopBack[T any](l *List[T]) {
	require(l.size > 0, "list cannot be empty")
	_ = Delete(RIter(l))
}

func PopFront[T any](l *List[T]) {
	require(l.size > 0, "list cannot be empty")
	_ = Delete(Iter(l))
}

// Front returns the first element in the list. It panics if the given list is
// empty.
// This function runs is O(1).
func Front[T any](l *List[T]) T {
	require(l.size > 0, "list cannot be empty")
	return l.head.next.value
}

// Back returns the last element in the list. It panics if the given list is
// empty.
// This function is O(1).
func Back[T any](l *List[T]) T {
	require(l.size > 0, "list cannot be empty")
	return l.tail.prev.value
}

func Reverse[T any](l *List[T]) {
	internal.Reverse[T](Iter(l), RIter(l), l.size)
}

func (l *List[T]) insertBetween(node, prev, next *node[T]) {
	node.next = next
	node.prev = prev

	next.prev = node
	prev.next = node

	l.size += 1
}

// Insert inserts given values just after the given iterator.
// This function is O(len(elems)). So inserting a single element would
// be O(1).
func Insert[T any](it Iterator[T], elems ...T) {
	switch typedIt := it.(type) {
	case *FrwIter[T]:
		require(typedIt.Valid(), "iterator must be valid")
		require(typedIt.node.next != nil, "bad iterator")

		for i := len(elems) - 1; i >= 0; i-- {
			elem := elems[i]
			node := &node[T]{value: elem}
			typedIt.lst.insertBetween(node, typedIt.node, typedIt.node.next)
		}
	case *RevIter[T]:
		require(typedIt.Valid(), "iterator must be valid")
		require(typedIt.node.prev != nil, "bad iterator")

		for _, elem := range elems {
			node := &node[T]{value: elem}
			typedIt.lst.insertBetween(node, typedIt.node.prev, typedIt.node)
		}
	default:
		panic("wrong iter type")
	}
}

func (l *List[T]) deleteNode(node *node[T]) (*node[T], *node[T]) {
	prev := node.prev
	next := node.next

	prev.next = next
	next.prev = prev

	node = nil

	l.size -= 1

	return prev, next
}

// Delete deletes the element that the given iterator `it` is pointing to.
// The given iterator be invalidated and cannot be used anymore. Instead you
// should use the returned iterator. The returned iterator points to the `it`'s
// previous element and its next element would be `it.Next()`.
// This function is O(1).
func Delete[T any](it Iterator[T]) Iterator[T] {
	switch typedIt := it.(type) {
	case *FrwIter[T]:
		require(typedIt.Valid(), "iterator must be valid")
		require(typedIt.node.prev != nil && typedIt.node.next != nil,
			"bad iterator")

		prev, _ := typedIt.lst.deleteNode(typedIt.node)
		return &FrwIter[T]{
			node: prev,
			lst:  typedIt.lst,
		}
	case *RevIter[T]:
		require(typedIt.Valid(), "iterator must be valid")
		require(typedIt.node.prev != nil && typedIt.node.next != nil,
			"bad iterator")

		_, next := typedIt.lst.deleteNode(typedIt.node)
		return &RevIter[T]{
			node: next,
			lst:  typedIt.lst,
		}
	default:
		panic("wrong iter type")
	}
}
