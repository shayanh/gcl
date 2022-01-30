package main

type Node[T any] struct {
	Value T
	Next  *Node[T]
	Prev  *Node[T]
}

type ListIter[T any] struct {
	Node *Node[T]
}

func (it *ListIter[T]) Next() {
	it.Node = it.Node.Next
}

func (it *ListIter[T]) Prev() {
	it.Node = it.Node.Prev
}

func (it *ListIter[T]) Done() bool {
	return it.Node == nil
}

func (it *ListIter[T]) Value() T {
	return it.Node.Value
}

func (it *ListIter[T]) SetValue(val T) {
	it.Node.Value = val
}

type ListRevIter[T any] struct {
	Node *Node[T]
}

func (it *ListRevIter[T]) Next() {
	it.Node = it.Node.Prev
}

func (it *ListRevIter[T]) Prev() {
	it.Node = it.Node.Next
}

func (it *ListRevIter[T]) Done() bool {
	return it.Node == nil
}

func (it *ListRevIter[T]) Value() T {
	return it.Node.Value
}

func (it *ListRevIter[T]) SetValue(val T) {
	it.Node.Value = val
}

type List[T any] struct {
	Front *Node[T]
	Back *Node[T]
	size int
}

func NewList[T any]() *List[T] {
	l := &List[T]{}
	return l
}

func(l *List[T]) PushBack(t T) {
	node := &Node[T]{
		Value: t,
		Next: nil,
		Prev: l.Back,
	}
	if l.Back != nil {
		l.Back.Next = node
	}
	l.Back = node
	if l.Front == nil {
		l.Front = node
	}
	l.size += 1
}

func (l *List[T]) Iter() *ListIter[T] {
	return &ListIter[T]{Node: l.Front}
}

func (l *List[T]) RevIter() *ListRevIter[T] {
	return &ListRevIter[T]{Node: l.Back}
}

func (l *List[T]) ForEach(fn func(T)) {
	forEach[T](l.Iter(), fn)
}

func (l *List[T]) ForEachIter(it Iter[T], fn func(T)) {
	forEach(it, fn)
}

func (l *List[T]) Size() int {
	return l.size
}

func (l *List[T]) Reverse() {
	reverse[T](l.Iter(), l.RevIter(), l.size)
}
