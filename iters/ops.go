package iters

import (
	"golang.org/x/exp/constraints"

	"github.com/shayanh/gcl"
)

// ForEach calls a function on each element of an iterator. ForEach moves the
// given iterator it to its end such that after a ForEach call it.HasNext() will
// be false.
func ForEach[T any](it Iterator[T], fn func(T)) {
	for it.HasNext() {
		fn(it.Next())
	}
}

// Equal determines if the elements of two iterators are equal. It returns true
// if all elements are equal and both iterators have the same number of
// elements. If the first iterator has l1 elements and the second iterator has
// l2 elements, Equal advances each iterator min(l1, l2) steps.
func Equal[T comparable](it1, it2 Iterator[T]) bool {
	for it1.HasNext() && it2.HasNext() {
		v1 := it1.Next()
		v2 := it2.Next()
		if v1 != v2 {
			return false
		}
	}
	if it1.HasNext() || it2.HasNext() {
		return false
	}
	return true
}

// EqualFunc works the same as Equal, but it uses the function eq for element
// comparison.
func EqualFunc[T1 any, T2 any](it1 Iterator[T1], it2 Iterator[T2], eq gcl.EqualFn[T1, T2]) bool {
	for it1.HasNext() && it2.HasNext() {
		v1 := it1.Next()
		v2 := it2.Next()
		if !eq(v1, v2) {
			return false
		}
	}
	if it1.HasNext() || it2.HasNext() {
		return false
	}
	return true
}

type mapIter[T any, V any] struct {
	wrappedIt Iterator[T]
	fn        func(T) V
}

func (it *mapIter[T, V]) HasNext() bool {
	return it.wrappedIt.HasNext()
}

func (it *mapIter[T, V]) Next() V {
	return it.fn(it.wrappedIt.Next())
}

// Map applies the function fn on elements the given iterator it and returns a
// new iterator over the mapped values. Map moves the given iterator it to its
// end such that after a Map call it.HasNext() will be false.
// Map is lazy, in a way that if you don't consume the returned iterator
// nothing will happen.
func Map[T any, V any](it Iterator[T], fn func(T) V) Iterator[V] {
	return &mapIter[T, V]{wrappedIt: it, fn: fn}
}

// Reduce applies a function of two arguments cumulatively to the items of
// the given iterator from the beginning to the end.
// Reduce moves the given iterator it to its end such that after a Reduce
// call, it.HasNext() will be false.
func Reduce[T any](it Iterator[T], fn func(T, T) T) (acc T) {
	if !it.HasNext() {
		return
	}
	acc = it.Next()
	for it.HasNext() {
		acc = fn(acc, it.Next())
	}
	return
}

// Fold applies a function of two arguments cumulatively to the items of
// the given iterator from the beginning to the end. Fold gives an initial
// value and start its operation by using the initial value.
// Fold moves the given iterator it to its end such that after a Fold
// call it.HasNext() will be false.
func Fold[T any, V any](it Iterator[T], fn func(V, T) V, init V) (acc V) {
	acc = init
	for it.HasNext() {
		acc = fn(acc, it.Next())
	}
	return
}

type nextState int

const (
	unknown nextState = iota
	hasNext
	noNext
)

type filterIter[T any] struct {
	wrappedIt Iterator[T]
	pred      func(T) bool
	state     nextState
	next      T
}

func (it *filterIter[T]) findNext() {
	for it.wrappedIt.HasNext() {
		v := it.wrappedIt.Next()
		if it.pred(v) {
			it.state = hasNext
			it.next = v
			return
		}
	}
	it.state = noNext
}

func (it *filterIter[T]) HasNext() bool {
	if it.state == unknown {
		it.findNext()
	}
	if it.state == hasNext {
		return true
	}
	return false
}

func (it *filterIter[T]) Next() T {
	if it.state == unknown {
		it.findNext()
	}
	if it.state == noNext {
		panic("iterator does not have next")
	}
	it.state = unknown
	return it.next
}

// Filter filters elements of an iterator that satisfy the pred.
// Filter returns an iterator over the filtered elements. Filter moves the given
// iterator it to its end such that after a Filter call it.HasNext() will be
// false.
// Filter is lazy, in a way that if you don't consume the returned iterator
// nothing will happen.
func Filter[T any](it Iterator[T], pred func(T) bool) Iterator[T] {
	return &filterIter[T]{wrappedIt: it, pred: pred}
}

// Find returns the first element in an iterator that satisfies pred. The
// returned boolean value indicates if such an element exists. If a satisfying
// element exists, the given iterator it advances the first satisfying element
// by one step, otherwise Find moves the iterator it to its end.
func Find[T any](it Iterator[T], pred func(T) bool) (t T, ok bool) {
	ok = false
	for it.HasNext() {
		v := it.Next()
		if pred(v) {
			t, ok = v, true
			return
		}
	}
	return
}

// Max returns the maximum element in an iterator of any ordered type.
// Max moves the given iterator it to its end such that after a Max
// call it.HasNext() will be false.
func Max[T constraints.Ordered](it Iterator[T]) (max T) {
	if !it.HasNext() {
		return
	}
	max = it.Next()
	for it.HasNext() {
		v := it.Next()
		if max < v {
			max = v
		}
	}
	return
}

// MaxFunc returns the maximum element in an iterator and uses the given
// less function for comparison.
// MaxFunc moves the given iterator it to its end such that after a MaxFunc
// call it.HasNext() will be false.
func MaxFunc[T any](it Iterator[T], less gcl.LessFn[T]) (max T) {
	if !it.HasNext() {
		return
	}
	max = it.Next()
	for it.HasNext() {
		v := it.Next()
		if less(max, v) {
			max = v
		}
	}
	return
}

// Min returns the minimum element in an iterator of any ordered type.
// Min moves the given iterator it to its end such that after a Min
// call it.HasNext() will be false.
func Min[T constraints.Ordered](it Iterator[T]) (min T) {
	if !it.HasNext() {
		return
	}
	min = it.Next()
	for it.HasNext() {
		v := it.Next()
		if min > v {
			min = v
		}
	}
	return
}

// MinFunc returns the minimum element in an iterator and uses the given
// less function for comparison.
// MinFunc moves the given iterator it to its end such that after a MinFunc
// call it.HasNext() will be false.
func MinFunc[T constraints.Ordered](it Iterator[T], less gcl.LessFn[T]) (min T) {
	if !it.HasNext() {
		return
	}
	min = it.Next()
	for it.HasNext() {
		v := it.Next()
		if less(v, min) {
			min = v
		}
	}
	return
}

// Sum returns sum of the elements in an iterator of any numeric type.
// Sum moves the given iterator it to its end such that after a Sum
// call it.HasNext() will be false.
func Sum[T gcl.Number](it Iterator[T]) (sum T) {
	sum = 0
	for it.HasNext() {
		sum += it.Next()
	}
	return
}

type zipIter[T1 any, T2 any] struct {
	it1 Iterator[T1]
	it2 Iterator[T2]
}

func (it *zipIter[T1, T2]) HasNext() bool {
	return it.it1.HasNext() && it.it2.HasNext()
}

func (it *zipIter[T1, T2]) Next() gcl.Zipped[T1, T2] {
	return gcl.Zipped[T1, T2]{
		First:  it.it1.Next(),
		Second: it.it2.Next(),
	}
}

// Zip zips the two given iterators and returns a single iterator over
// gcl.Zipped values.
func Zip[T1 any, T2 any](it1 Iterator[T1], it2 Iterator[T2]) Iterator[gcl.Zipped[T1, T2]] {
	return &zipIter[T1, T2]{
		it1: it1,
		it2: it2,
	}
}
