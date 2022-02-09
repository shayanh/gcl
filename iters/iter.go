package iters

// We don't have a struct like container ops to provide all the algorithms (and
// then embed that struct into containers) because go doesn't support type
// variance and we cannot use custom iter for container.

type Iter[T any] interface {
	Next()
	Done() bool
	Value() T
}

type BiDrecIter[T any] interface {
	Iter[T]
	Prev()
}

type MutIter[T any] interface {
	Iter[T]
	SetValue(T)
}
