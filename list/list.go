package list

import (
	"github.com/shayanh/gogl"
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

	for _, elem := range elems {
		l.PushBack(elem)
	}
	return l
}

func (l *List[T]) Size() int {
	return l.size
}

// Iter returns a forward iterator to the beginning.
func (l *List[T]) Iter() Iter[T] {
	return &FrwIter[T]{
		node: l.head.next,
		lst:  l,
	}
}

// RIter returns a reverse iterator to the beginning (in the reverse order).
func (l *List[T]) RIter() Iter[T] {
	return &RevIter[T]{
		node: l.tail.prev,
		lst:  l,
	}
}

func (l *List[T]) PushBack(t T) {
	_ = l.Insert(l.RIter(), t)
}

func (l *List[T]) PushFront(t T) {
	_ = l.Insert(l.Iter(), t)
}

func (l *List[T]) PopBack() {
	_ = l.Erase(l.RIter())
}

func (l *List[T]) PopFront() {
	_ = l.Erase(l.Iter())
}

func (l *List[T]) ForEach(fn func(T)) {
	internal.ForEach[T](l.Iter(), fn)
}

func (l *List[T]) Reverse() {
	internal.Reverse[T](l.Iter(), l.RIter(), l.size)
}

func (l *List[T]) insertBetween(node, prev, next *node[T]) {
	node.next = next
	node.prev = prev

	next.prev = node
	prev.next = node

	l.size += 1
}

func (l *List[T]) Insert(it Iter[T], elems ...T) Iter[T] {
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

func (l *List[T]) Erase(it Iter[T]) Iter[T] {
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

func (l *List[T]) MaxFunc(less gogl.LessFn[T]) T {
	return internal.MaxFunc[T](l.Iter(), less)
}

func (l *List[T]) MinFunc(less gogl.LessFn[T]) T {
	return internal.MinFunc[T](l.Iter(), less)
}
