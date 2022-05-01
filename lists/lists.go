// Package lists provides a doubly linked list and various functions useful
// with lists of any type.
package lists

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/shayanh/gcl"
	"github.com/shayanh/gcl/internal"
	"github.com/shayanh/gcl/iters"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
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

// New creates a new linked list and returns a pointer to it.
func New[T any](elems ...T) *List[T] {
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
	l := New[T]()
	for it.HasNext() {
		PushBack(l, it.Next())
	}
	return l
}

func (l *List[T]) String() string {
	var b strings.Builder
	b.WriteString("lists.List[")
	it := Iter(l)
	for it.HasNext() {
		v := it.Next()
		fmt.Fprintf(&b, "%v", v)
		if it.HasNext() {
			b.WriteString(" ")
		}
	}
	b.WriteString("]")
	return b.String()
}

func (l *List[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(toSlice(l))
}

func (l *List[T]) UnmarshalJSON(b []byte) error {
	var slice []T
	if err := json.Unmarshal(b, &slice); err != nil {
		return err
	}
	*l = *New(slice...)
	return nil
}

// Len returns size of the given list. This function is O(1).
func Len[T any](l *List[T]) int {
	return l.size
}

// Iter returns a forward iterator to the beginning. Initially, the returned
// iterator is located at one step before the first element (one-before-first).
func Iter[T any](l *List[T]) Iterator[T] {
	return &FrwIter[T]{
		node: l.head,
		lst:  l,
	}
}

// RIter returns a reverse iterator going from the end to the beginning.
// Initially, the returned iterator is located at one step past the last
// element (one-past-last).
func RIter[T any](l *List[T]) Iterator[T] {
	return &RevIter[T]{
		node: l.tail,
		lst:  l,
	}
}

// Equal tests whether two lists are equal: the same length and all elements
// equal. Floating point NaNs are not considered equal.
// This function is O(min(Len(l1), Len(l2))).
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

// EqualFunc tests whether two lists are equal using the given `eq` function.
// For each pair of elements, `eq` determines if they are equal or not.
// This function is O(f * min(Len(l1), Len(l2))), where f is the time complexity
// of `eq` function.
func EqualFunc[T1 any, T2 any](l1 *List[T1], l2 *List[T2], eq gcl.EqualFn[T1, T2]) bool {
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
// The elements are compared sequentially from the beginning, until one element
// is not equal to the corresponding one.
// The result of comparing the first non-matching elements is returned. If
// both lists are equal until one of them ends, the shorter list is considered
// less than the longer one.
// The result is 0 if l1 == l2, -1 if l1 < l2, and +1 if l1 > l2.
// Comparisons involving floating point NaNs are ignored.
// This function is O(min(Len(l1), Len(l2))).
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

// CompareFunc operates the same as Compare function but it uses the given
// `cmp` function for comparing each pair of elements.
// The result is the first non-zero result of cmp; if cmp always
// returns 0 the result is 0 if Len(l1) == Len(l2), -1 if Len(l1) < Len(l2),
// This function is O(f * min(Len(l1), Len(l2))), where f is the time complexity
// of `cmp` function.
func CompareFunc[T1 any, T2 any](l1 *List[T1], l2 *List[T2], cmp gcl.CompareFn[T1, T2]) int {
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

// PushBack appends the given elements to the back of list `l`.
// This function is O(Len(elems)). So for a single element it would be O(1).
func PushBack[T any](l *List[T], elems ...T) {
	Insert(RIter(l), elems...)
}

// PushFront appends the given elements to the beginning of list `l`.
// This function is O(Len(elems)). So for a single element it would be O(1).
func PushFront[T any](l *List[T], elems ...T) {
	Insert(Iter(l), elems...)
}

// PopBack deletes the last element in the list. It requires the list to be
// non-empty, otherwise it panics.
// This function is O(1).
func PopBack[T any](l *List[T]) {
	require(l.size > 0, "list cannot be empty")
	_ = Delete(RIter(l))
}

// PopFront deletes the last element in the list. It requires the list to be
// non-empty, otherwise it panics.
// This function is O(1).
func PopFront[T any](l *List[T]) {
	require(l.size > 0, "list cannot be empty")
	_ = Delete(Iter(l))
}

// Front returns the first element in the list. It panics if the given list is
// empty.
// This function is O(1).
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

// Reverse reverses the elements of the given list.
// This function is O(n), where n is length of the list.
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

// Insert inserts the given values next after the given iterator. Direction of
// the next is determined by the given iterator type. Insert panics if the given
// iterator is invalid.
// This function is O(Len(elems)). So inserting a single element would be O(1).
func Insert[T any](it Iterator[T], elems ...T) {
	require(it.Valid(), "iterator must be valid")
	switch typedIt := it.(type) {
	case *FrwIter[T]:
		require(typedIt.node.next != nil, "bad iterator")
		for i := len(elems) - 1; i >= 0; i-- {
			elem := elems[i]
			node := &node[T]{value: elem}
			typedIt.lst.insertBetween(node, typedIt.node, typedIt.node.next)
		}
	case *RevIter[T]:
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

// Delete deletes the element that the given iterator `it` is pointing to it. The
// iterator `it` will be invalidated and cannot be used anymore. You should use
// the returned iterator instead. The returned iterator points to `it`'s
// previous element and its next element is `it.Next()`.
// Delete requires `it` to points to an actual element in the list. For example,
// it doesn't accept Iter and RIter iterators (in their inital state) because
// these iterators point to one-before-first and one-past-last respectively and
// these are not actual list elements. Also, the given iterator must be valid.
// Delete panics if `it` doesn't have the above requirements.
// This function is O(1).
func Delete[T any](it Iterator[T]) Iterator[T] {
	require(it.Valid(), "iterator must be valid")
	switch typedIt := it.(type) {
	case *FrwIter[T]:
		require(typedIt.node.prev != nil && typedIt.node.next != nil,
			"bad iterator")
		prev, _ := typedIt.lst.deleteNode(typedIt.node)
		return &FrwIter[T]{
			node: prev,
			lst:  typedIt.lst,
		}
	case *RevIter[T]:
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

// Sort sorts a list of any ordered type in ascending order.
// This function is O(n * log(n)), where n is length of the list.
func Sort[T constraints.Ordered](l *List[T]) {
	slice := toSlice(l)
	slices.Sort(slice)
	it := Iter(l)
	for _, v := range slice {
		it.Next()
		it.Set(v)
	}
}

// SortFunc sorts the given list in ascending order as determined by the `less`
// function.
// This function is O(f * n * log(n)), where n is length of the list and
// f is time complexity of the `less` function.
func SortFunc[T any](l *List[T], less gcl.LessFn[T]) {
	slice := toSlice(l)
	slices.SortFunc(slice, less)
	it := Iter(l)
	for _, v := range slice {
		it.Next()
		it.Set(v)
	}
}

// IsSorted tests whether a list of any ordered type is sorted in ascending order.
// This function is O(n), where n is length of the list.
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

// IsSortedFunc tests whether a list type is sorted in ascending order, according
// to the `less` comparison function.
// This function is O(f * n), where n is length of the list and f is time
// complexity of the `less` function.
func IsSortedFunc[T any](l *List[T], less gcl.LessFn[T]) bool {
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

// Compact replaces every consecutive group of equal elements with a single
// copy. This is like the uniq Unix command.
// This function is O(n), where n is length of the list.
func Compact[T comparable](l *List[T]) {
	it1 := Iter(l).(*FrwIter[T])
	if !it1.HasNext() {
		return
	}
	last := it1.Next()
	newSize := 1
	if !it1.HasNext() {
		return
	}
	it1.Next()

	it2 := Iter(l)
	it2.Next()
	for it2.HasNext() {
		v := it2.Next()
		if v != last {
			it1.Set(v)
			if it1.HasNext() {
				it1.Next()
			}
			last = v
			newSize += 1
		}
	}
	it1.lst.tail.prev = it1.node.prev
	it1.node.prev.next = it1.lst.tail

	it1.lst.size = newSize
}

// CompactFunc is like Compact but it uses the `eq` function for comparison.
// This function is O(f * n), where n is length of the list and f is time
// complexity of the `eq` function.
func CompactFunc[T any](l *List[T], eq gcl.EqualFn[T, T]) {
	it1 := Iter(l).(*FrwIter[T])
	if !it1.HasNext() {
		return
	}
	last := it1.Next()
	newSize := 1
	if !it1.HasNext() {
		return
	}
	it1.Next()

	it2 := Iter(l)
	it2.Next()
	for it2.HasNext() {
		v := it2.Next()
		if !eq(v, last) {
			it1.Set(v)
			if it1.HasNext() {
				it1.Next()
			}
			last = v
			newSize += 1
		}
	}
	it1.lst.tail.prev = it1.node.prev
	it1.node.prev.next = it1.lst.tail

	it1.lst.size = newSize
}

// Index returns the index of the first occurrence of v in l, or -1 if not
// present.
// This function is O(n), where n is length of the list.
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

// IndexFunc returns the index of the first element satisfying pred,
// or -1 if none do.
// This function is O(f * n), where n is length of the list and f is time
// complexity of the `pred` function.
func IndexFunc[T comparable](l *List[T], pred func(T) bool) int {
	i := 0
	it := Iter(l)
	for it.HasNext() {
		v := it.Next()
		if pred(v) {
			return i
		}
		i++
	}
	return -1
}

// Pos returns an iterator pointing to the first occurrence of v in l.
// The returned boolean value indicates if v is present in the list.
// This function is O(n), where n is length of the list.
func Pos[T comparable](l *List[T], v T) (Iterator[T], bool) {
	it := Iter(l)
	for it.HasNext() {
		if it.Next() == v {
			return it, true
		}
	}
	return it, false
}

// Contains tests whether the given list l has value v.
// This function is O(n), where n is length of the list.
func Contains[T comparable](l *List[T], v T) bool {
	_, res := Pos(l, v)
	return res
}

// Clone returns a copy of the given list. The elements are copied using
// assignment, so this is a shallow clone.
// This function is O(n), where n is length of the list.
func Clone[T any](l *List[T]) *List[T] {
	res := New[T]()
	it := Iter(l)
	for it.HasNext() {
		PushBack(res, it.Next())
	}
	return res
}
