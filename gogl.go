package gogl

import "golang.org/x/exp/constraints"

type LessFn[T any] func(a, b T) bool

func Less[T constraints.Ordered](a, b T) bool {
	return a < b
}

type EqualFn[T any] func(a, b T) bool

func Equal[T comparable](a, b T) bool {
	return a == b
}

type CompareFn[T1 any, T2 any] func(T1, T2) int

type Number interface {
	constraints.Integer | constraints.Float | constraints.Complex
}
