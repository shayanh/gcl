package list

import (
	"github.com/shayanh/gogl/internal"
)

type node[T any] struct {
	value T
	next  *node[T]
	prev  *node[T]
}

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

func Len[T any](l *List[T]) int {
	return l.size
}


// Begin returns a forward iterator to the beginning.
func Begin[T any](l *List[T]) Iter[T] {
	return &FrwIter[T]{
		node: l.head.next,
		lst:  l,
	}
}

// RBegin returns a reverse iterator to the beginning (in the reverse order).
func RBegin[T any](l *List[T]) Iter[T] {
	return &RevIter[T]{
		node: l.tail.prev,
		lst:  l,
	}
}

func PushBack[T any](l *List[T], elems ...T) {
	_ = Insert(l, RBegin(l), elems...)
}

func PushFront[T any](l *List[T], elems ...T) {
	_ = Insert(l, Begin(l), elems...)
}

func PopBack[T any](l *List[T]) {
	_ = Erase(l, RBegin(l))
}

func PopFront[T any](l *List[T]) {
	_ = Erase(l, Begin(l))
}

func Reverse[T any](l *List[T]) {
	internal.Reverse[T](Begin(l), RBegin(l), l.size)
}

func (l *List[T]) insertBetween(node, prev, next *node[T]) {
	node.next = next
	node.prev = prev

	next.prev = node
	prev.next = node

	l.size += 1
}

func Insert[T any](l *List[T], it Iter[T], elems ...T) Iter[T] {
	switch typedIt := it.(type) {
	case *FrwIter[T]:
		if typedIt.lst != l {
			panic("iterator doesn't belong to this list")
		}
		rit := typedIt.clone()
		for i, elem := range elems {
			node := &node[T]{value: elem}
			l.insertBetween(node, typedIt.node.prev, typedIt.node)
			if i == 0 {
				rit = &FrwIter[T]{
					node: node,
					lst:  l,
				}
			}
		}
		return rit
	case *RevIter[T]:
		if typedIt.lst != l {
			panic("iterator doesn't belong to this list")
		}
		rit := typedIt.clone()
		for i := len(elems) - 1; i >= 0; i-- {
			elem := elems[i]
			node := &node[T]{value: elem}
			l.insertBetween(node, typedIt.node, typedIt.node.next)
			if i == 0 {
				rit = &RevIter[T]{
					node: node,
					lst:  l,
				}
			}
		}
		return rit
	default:
		panic("wrong iter type")
	}
}

func (l *List[T]) erase(node *node[T]) (*node[T], *node[T]) {
	prev := node.prev
	next := node.next

	prev.next = next
	next.prev = prev

	node = nil

	return prev, next
}

func Erase[T any](l *List[T], it Iter[T]) Iter[T] {
	switch typedIt := it.(type) {
	case *FrwIter[T]:
		if typedIt.lst != l {
			panic("iterator doesn't belong to this list")
		}
		_, next := l.erase(typedIt.node)
		return &FrwIter[T]{
			node: next,
			lst:  l,
		}
	case *RevIter[T]:
		if typedIt.lst != l {
			panic("iterator doesn't belong to this list")
		}
		prev, _ := l.erase(typedIt.node)
		return &RevIter[T]{
			node: prev,
			lst:  l,
		}
	default:
		panic("wrong iter type")
	}
}
