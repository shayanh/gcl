// Package list provides a doubly linked list.
package lists

import (
	"sort"

	"github.com/shayanh/gogl"
	"github.com/shayanh/gogl/internal"
	"github.com/shayanh/gogl/iters"
	"golang.org/x/exp/constraints"
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

// Compare compares the elements of l1 and l2.
func Compare[T constraints.Ordered](l1, l2 *List[T]) int {
	it1, it2 := Iter(l1), Iter(l2)
	for it1.HasNext() {
		if !it2.HasNext() {
			return +1
		}
		v1, v2 := it1.Next(), it2.Next()
		switch {
		case v1 < v2:
			return -1
		case v1 > v2:
			return +1
		}
	}
	if it2.HasNext() {
		return -1
	}
	return 0
}

func CompareFunc[T1 any, T2 any](l1 *List[T1], l2 *List[T2], cmp gogl.CompareFn[T1, T2]) int {
	it1, it2 := Iter(l1), Iter(l2)
	for it1.HasNext() {
		if !it2.HasNext() {
			return +1
		}
		v1, v2 := it1.Next(), it2.Next()
		if c := cmp(v1, v2); c != 0 {
			return c
		}
	}
	if it2.HasNext() {
		return -1
	}
	return 0

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

func toSlice[T any](l *List[T]) []T {
	var res []T
	for it := Iter(l); it.HasNext(); {
		res = append(res, it.Next())
	}
	return res
}

func Sort[T constraints.Ordered](l *List[T]) {
	slice := toSlice(l)
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	it := Iter(l)
	for _, v := range slice {
		it.Next()
		it.Set(v)
	}
}

func SortFunc[T any](l *List[T], less gogl.LessFn[T]) {
	slice := toSlice(l)
	sort.Slice(slice, func(i, j int) bool {
		return less(slice[i], slice[j])
	})
	it := Iter(l)
	for _, v := range slice {
		it.Next()
		it.Set(v)
	}
}

func IsSorted[T constraints.Ordered](l *List[T]) bool {
	it := RIter(l)
	if !it.HasNext() {
		return true
	}
	vr := it.Next()
	for it.HasNext() {
		vl := it.Next()
		if vr < vl {
			return false
		}
		vr = vl
	}
	return true
}

func IsSortedFunc[T any](l *List[T], less gogl.LessFn[T]) bool {
	it := RIter(l)
	if !it.HasNext() {
		return true
	}
	vr := it.Next()
	for it.HasNext() {
		vl := it.Next()
		if less(vr, vl) {
			return false
		}
		vr = vl
	}
	return true
}

func Compact[T comparable](l *List[T]) {
	panic("impelement me")
}

func CompactFunc[T any](l *List[T], eq gogl.EqualFn[T]) {
	panic("implement me")
}

func Index[T comparable](l *List[T], v T) int {
	i := 0
	it := Iter(l)
	for it.HasNext() {
		if it.Next() == v {
			return i
		}
		i++
	}
	return -1
}

func IndexFunc[T comparable](l *List[T], v T, eq gogl.EqualFn[T]) int {
	i := 0
	it := Iter(l)
	for it.HasNext() {
		if eq(it.Next(), v) {
			return i
		}
		i++
	}
	return -1
}

func Pos[T comparable](l *List[T], v T) (Iterator[T], bool) {
	it := Iter(l)
	for it.HasNext() {
		if it.Next() == v {
			return it, true
		}
	}
	return it, false
}

func PosFunc[T any](l *List[T], v T, eq gogl.EqualFn[T]) (Iterator[T], bool) {
	it := Iter(l)
	for it.HasNext() {
		if eq(it.Next(), v) {
			return it, true
		}
	}
	return it, false
}

func Contains[T comparable](l *List[T], v T) bool {
	_, res := Pos(l, v)
	return res
}

func Clone[T any](l *List[T]) *List[T] {
	res := NewList[T]()
	it := Iter(l)
	for it.HasNext() {
		PushBack(res, it.Next())
	}
	return res
}
