package iter

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

type Iterable[T any] interface {
	Iter() Iter[T]
	RIter() Iter[T]
}
