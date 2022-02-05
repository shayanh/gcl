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

type MutIterableIterableAdapter[T any] struct {
	mut MutIterable[T]
}

func (adapter MutIterableIterableAdapter[T]) Iter() Iter[T] {
	return adapter.mut.MIter()
}

func (adapter MutIterableIterableAdapter[T]) RIter() Iter[T] {
	return adapter.mut.MRIter()
}

type MutIterable[T any] interface {
	MIter() MutIter[T]
	MRIter() MutIter[T]
}

func MutIterableAsIterable[T any](mutIterable MutIterable[T]) Iterable[T] {
	return MutIterableIterableAdapter[T]{
		mut: mutIterable,
	}
}
