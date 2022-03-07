package gogl

import "golang.org/x/exp/constraints"

type LessFn[T any] func(a, b T) bool

func Less[T constraints.Ordered](a, b T) bool {
	return a < b
}

func Greater[T constraints.Ordered](a, b T) bool {
	return a > b
}

type EqualFn[T1 any, T2 any] func(a T1, b T2) bool

func Equal[T comparable](a, b T) bool {
	return a == b
}

type CompareFn[T1 any, T2 any] func(T1, T2) int

func Compare[T constraints.Ordered](a, b T) int {
	if a < b {
		return -1
	}
	if a == b {
		return 0
	}
	return 1
}

type Number interface {
	constraints.Integer | constraints.Float | constraints.Complex
}
