# Goals

1. Simple API.
2. User should not be able to do something wrong with the API.
3. For any task there should be only one way to do it.

# Discussion

We decided to provide functions in packages instead of methods, since methods
cannot have extra generic types. Also with functions, we would have a similar
API to the standard Go library.

`iters` package is inspired by Rust’s iterators. However, in general we won’t be
as strict as Rust and we will allow some extra operations on collections that
Rust doesn’t.

# Packages

## iters

Package iters defines the general iterator interface and provides operations on
the given iterators.

```go
type Iterator[T any] interface {
	HasNext() bool
	Next() T
}

func Advance(Iter[T], n int)

func ForEach(Iter[T], fn)

func Map(Iter[T], func(T) V) []V

func Filter(Iter[T], func(T) bool) []V

func Reduce(Iter[T], func(T, T) T) V
func Fold(Iter[T], func(V, T) V, V) V

func Sum(Iter[T]) Integer

func Max(Iter[T]) T
func Min(Iter[T]) T

func MaxFunc(Iter[T], lessFn) T
func MinFunc(Iter[T], lessFn) T

func Merge(it1, it2 Iter[T]) []T
func MergeFunc(it1, it2 Iter[T], lessFn) []T
```

## set

Ordered Set

## ordmap

Ordered Map

## hashmap

[Unordered] Hash Map

## list

Package list provides a doubly linked list.

```go
type List[T] struct

func NewList[T](elems ...T) *List[T]

func Len(l) int

func ّIter(l) Iter[T]
func RIter(l) Iter[T]

func Equal(l1, l2 *List[T]) bool
func EqualFunc(l1, l2 *List[T], eqFn) bool

func PushBack(l, ...T)
func PushFront(l, ...T)

func PopBack(l)
func PopFront(l)

func Back(l) T
func Front(l) T

func Insert(Iter[T], ...T)
func Delete(Iter[T]) Iter[T]

func Reverse(l)

func Sort(l)
func SortFunc(l, lessFn)

func Compact(l)
func CompactFunc(l, lessFn)

func MaxElem(l) Iter[T]
func MaxElemFunc(l, lessFn) Iter[T]

func MinElem(l) Iter[T]
func MinElemFunc(l, lessFn) Iter[T]

func Find(l, T) Iter[T]
func FindFunc(l, T, lessFn) Iter[T]
func Contains(l, T) bool

func Clone() *List[T]
```

## mapops

Extra operations for builtin Go maps.

## sliceops

Extra operations for builtin Go slices.

## *internal*
