package main

type Iter[T any] interface {
	Next()
	Done() bool
	Value() T
}

type MutIter[T any] interface {
	Iter[T]
	SetValue(T)
}

func forEach[T any](it Iter[T], fn func(T)) {
	for ; !it.Done(); it.Next() {
		fn(it.Value())
	}
}

func reverse[T any](fIt MutIter[T], rIt MutIter[T], length int) {
	fIdx, rIdx := 0, length - 1
	for fIdx < rIdx {
		fVal, rVal := fIt.Value(), rIt.Value()
		fIt.SetValue(rVal)
		rIt.SetValue(fVal)

		fIdx += 1
		rIdx -= 1
		fIt.Next()
		rIt.Next()
	}
}
