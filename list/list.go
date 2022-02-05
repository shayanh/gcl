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
	front *node[T]
	back  *node[T]
	size  int
}

func NewList[T any]() *List[T] {
	front := &node[T]{}
	back := &node[T]{}

	front.next = back
	back.prev = front
	l := &List[T]{
		front: front,
		back:  back,
	}
	return l
}

func (l *List[T]) Size() int {
	return l.size
}

// Is this better or having Begin(), End()?
func (l *List[T]) Iter() Iter[T] {
	return &FrwIter[T]{
		node: l.front.next,
		lst:  l,	
	}
}

// Is this better or having RBegin(), REnd()?
func (l *List[T]) RIter() Iter[T] {
	return &RevIter[T]{
		node: l.back.prev,
		lst:  l,
	}
}

// TODO: should we accept an array of values? 
func (l *List[T]) PushBack(t T) {
	l.Insert(l.RIter(), t)	
}

// TODO: should we accept an array of values? 
func (l *List[T]) PushFront(t T) {
	l.Insert(l.Iter(), t)
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

// TODO: should we accept an array of values?
// Should we return an iter here?
func (l *List[T]) Insert(it Iter[T], val T) {
	node := &node[T]{value: val}

	switch tip := it.(type) {
	case *FrwIter[T]:
		if tip.lst != l {
			panic("iterator doesn't belong to this list")
		}
		l.insertBetween(node, tip.node.prev, tip.node)
	case *RevIter[T]:
		if tip.lst != l {
			panic("iterator doesn't belong to this list")
		}
		l.insertBetween(node, tip.node, tip.node.next)
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
	switch tip := it.(type) {
	case *FrwIter[T]:
		if tip.lst != l {
			panic("iterator doesn't belong to this list")
		}
		_, next := l.erase(tip.node)
		it = nil
		return &FrwIter[T]{
			node: next,
			lst: l,
		}
	case *RevIter[T]:
		if tip.lst != l {
			panic("iterator doesn't belong to this list")
		}
		prev, _ := l.erase(tip.node)
		it = nil
		return &RevIter[T]{
			node: prev,
			lst: l,
		}
	default:
		panic("wrong iter type")
	}
}
