package list

import (
	"github.com/shayanh/gogl/internal"
	"github.com/shayanh/gogl/iter"
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

func (l *List[T]) Iter() *Iter[T] {
	return &Iter[T]{
		node: l.front.next,
		lst:  l,	
	}
}

func (l *List[T]) RIter() *RIter[T] {
	return &RIter[T]{
		node: l.back.prev,
		lst:  l,
	}
}

// TODO: accept an array of values
func (l *List[T]) PushBack(t T) {
	l.Insert(l.RIter(), t)	
}

func (l *List[T]) PushFront(t T) {
	l.Insert(l.Iter(), t)
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

// TODO: accept an array of values
func (l *List[T]) Insert(it iter.Iter[T], val T) {
	node := &node[T]{value: val}

	switch tip := it.(type) {
	case *Iter[T]:
		if tip.lst != l {
			panic("iterator doesn't belong to this list")
		}
		l.insertBetween(node, tip.node.prev, tip.node)
	case *RIter[T]:
		if tip.lst != l {
			panic("iterator doesn't belong to this list")
		}
		l.insertBetween(node, tip.node, tip.node.next)
	default:
		panic("wrong iter type")
	}
}

//func (l *List[T]) erase(node *node[T]) {
	
//}

//func (l *List[T]) Erase(it iter.Iter[T]) {
	//switch tip := it.(type) {
	//case *Iter[T]:
		//if tip.lst != l {
			//panic("iterator doesn't belong to this list")
		//}
		//l.insertBetween(node, tip.node.prev, tip.node)
	//case *RIter[T]:
		//if tip.lst != l {
			//panic("iterator doesn't belong to this list")
		//}
		//l.insertBetween(node, tip.node, tip.node.next)
	//default:
		//panic("wrong iter type")
	//}
//}
