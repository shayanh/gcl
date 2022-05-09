// Package gcl defines common data types.
package gcl

import "golang.org/x/exp/constraints"

// LessFn defines a less than function for pairs of any type.
type LessFn[T any] func(a, b T) bool

// Less is a function of type LessFn, which operates on the comparable types.
// Less(a, b) is true if a < b.
func Less[T constraints.Ordered](a, b T) bool {
	return a < b
}

// Greater is a function of type LessFn, which operates on the comparable types.
// Greater(a, b) is true if a > b.
func Greater[T constraints.Ordered](a, b T) bool {
	return a > b
}

// Equal defines a equal function for pairs of any two types.
type EqualFn[T1 any, T2 any] func(a T1, b T2) bool

// Equal is a function of type EqualFn, which operates on the comparable types.
func Equal[T comparable](a, b T) bool {
	return a == b
}

// Equal defines a compare function for pairs of any two types.
type CompareFn[T1 any, T2 any] func(T1, T2) int

// Compare is a function of types CompareFn, which operates of ordered types.
// Compare(a, b) equals to:
// -1 if a < b
// 0 if a == b
// +1 if a > b
func Compare[T constraints.Ordered](a, b T) int {
	if a < b {
		return -1
	}
	if a == b {
		return 0
	}
	return 1
}

// Number contains all the different numeric types.
type Number interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

// MapElem is an element in a map data type.
type MapElem[K any, V any] struct {
	Key   K
	Value V
}

// Zipped contains values of two different types. You can access the values
// through the First and Second fields.
type Zipped[T1 any, T2 any] struct {
	First  T1
	Second T2
}
