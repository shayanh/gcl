package gogl

import "constraints"

type LessFn[T any] func(a, b T) bool

func Less[T constraints.Ordered](a, b T) bool {
	return a < b
}

type EqualFn[T any] func (a, b T) bool

func Equal[T comparable](a, b T) bool {
	return a == b
}
