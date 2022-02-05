package internal

import (
	"github.com/shayanh/gogl/iter"
)

type ContainerOps[T any] struct {
	iterable iter.Iterable[T]
}

func MakeContainerOps[T any](iterable iter.Iterable[T]) ContainerOps[T] {
	return ContainerOps[T]{
		iterable: iterable,
	}
}

func (ops ContainerOps[T]) Iter() iter.Iter[T] {
	return ops.iterable.Iter()
}

func (ops ContainerOps[T]) RIter() iter.Iter[T] {
	return ops.iterable.RIter()
}

func (ops ContainerOps[T]) ForEach(fn func(T)) {
	for it := ops.iterable.Iter(); !it.Done(); it.Next() {
		fn(it.Value())
	}
}

type MutContainerOps[T any] struct {
	ContainerOps[T]
	mIterable iter.MutIterable[T]
}

func MakeMutContainerOps[T any](mIterable iter.MutIterable[T]) MutContainerOps[T] {
	return MutContainerOps[T]{
		ContainerOps: MakeContainerOps(iter.MutIterableAsIterable(mIterable)),
		mIterable:    mIterable,
	}
}

func (ops MutContainerOps[T]) Reverse(length int) {
	fIdx, rIdx := 0, length-1
	fIt, rIt := ops.mIterable.MIter(), ops.mIterable.MRIter()
	for fIdx < rIdx {
		fVal, rVal := fIt.Value(), rIt.Value()
		fIt.SetValue(rVal)
		rIt.SetValue(fVal)

		fIdx += 1
		rIdx -= 1
		fIt.Next()
		rIt.Next()

		if fIt.Done() || rIt.Done() {
			panic("bad iterator")
		}
	}
}
