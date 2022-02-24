# Goals

1. Simple and clear API. Behavior and time complexity of all functions should be
   documented clearly.
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

func ForEach(Iter[T], fn)

func Map(Iter[T], func(T) V) []V

func Filter(Iter[T], func(T) bool) []V

func Reduce(Iter[T], func(T, T) T) V
func Fold(Iter[T], func(V, T) V, V) V

func Find(Iter[T], func(T) bool) (T, bool)

func Sum(Iter[T]) Integer

func Max(Iter[T]) T
func Min(Iter[T]) T

func MaxFunc(Iter[T], lessFn) T
func MinFunc(Iter[T], lessFn) T

func Zip(it1 Iter[T], it2 Iter[V]) Iter[Zipped[T, V]]

// Not sure
func Advance(Iter[T], n int)

// Not sure
func Merge(it1, it2 Iter[T]) []T
func MergeFunc(it1, it2 Iter[T], lessFn) []T
```

## tsets

[Ordered] Tree Set

## hsets

[Unordered] Hash Set

## tmaps

[Ordered] Tree Map

## hmaps

[Unordered] Hash Map

## lists

Package lists provides a doubly linked list.

```go
type List[T] struct

func NewList[T](elems ...T) *List[T]

func FromIter(iters.Iter[T]) *List[T]

func Len(l) int

func ّIter(l) Iter[T]
func RIter(l) Iter[T]

func Equal(l1, l2 *List[T]) bool
func EqualFunc(l1, l2 *List[T], eqFn) bool

func Compare(l1, l2 *List[T]) int
func CompareFunc(l1, l2 *List[T], lessFn) int

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

func IsSorted(l) bool
func IsSortedFunc(l, lessFn) bool

func Compact(l)
func CompactFunc(l, lessFn)

func Index(l, T) int
func IndexFunc(l, T, eqFn) int

func Pos(l, T) (Iter[T], bool)
func PosFunc(l, T, eqFn) (Iter[T], bool)
func Contains(l, T) bool

func Clone() *List[T]
```

## gomaps

Extra operations for builtin Go maps.

```go

Iter()

FromIter()

```

## goslices

Extra operations for builtin Go slices.

```go

Iter()
RIter()

Reverse()

FromIter()

```

## *internal*
