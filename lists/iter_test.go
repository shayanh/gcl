package lists

import (
	"testing"

	"github.com/shayanh/gcl/iters"
)

func TestIterMut(t *testing.T) {
	l := New(1, 2, 3)
	it := IterMut(l)
	for it.HasNext() {
		v := it.Next()
		*v = *v + 1
	}
	want := New(2, 3, 4)
	if !Equal(l, want) {
		t.Error("wrong IterMut behavior")
	}
}

func TestRIterMut(t *testing.T) {
	l := New(1, 2, 3)
	it := RIterMut(l)
	for it.HasNext() {
		v := it.Next()
		*v = *v + 1
	}
	want := New(2, 3, 4)
	if !Equal(l, want) {
		t.Error("wrong IterMut behavior")
	}
}

var frwIterMutInsertTests = []struct {
	l     *List[int]
	itFn  func(*List[int]) *FrwIterMut[int]
	elems []int
	want  *List[int]
}{
	{
		l: New(1, 3),
		itFn: func(l *List[int]) *FrwIterMut[int] {
			it := IterMut(l)
			iters.Advance[*int](it, 1)
			return it
		},
		elems: []int{2},
		want:  New(1, 2, 3),
	},
	{
		l: New(1, 2, 5, 6),
		itFn: func(l *List[int]) *FrwIterMut[int] {
			it := IterMut(l)
			iters.Advance[*int](it, 2)
			return it
		},
		elems: []int{3, 4},
		want:  New(1, 2, 3, 4, 5, 6),
	},
}

func TestFrwIterMutInsert(t *testing.T) {
	for _, test := range frwIterMutInsertTests {
		it := test.itFn(test.l)
		if it.Insert(test.elems...); !Equal(test.l, test.want) {
			t.Errorf("Insert got %v, want %v", test.l, test.want)
		}
	}
}

var revIterMutInsertTests = []struct {
	l     *List[int]
	itFn  func(*List[int]) *RevIterMut[int]
	elems []int
	want  *List[int]
}{
	{
		l: New(1, 2, 5, 6),
		itFn: func(l *List[int]) *RevIterMut[int] {
			it := RIterMut(l)
			iters.Advance[*int](it, 2)
			return it
		},
		elems: []int{3, 4},
		want:  New(1, 2, 3, 4, 5, 6),
	},
}

func TestRevIterMutInsert(t *testing.T) {
	for _, test := range revIterMutInsertTests {
		it := test.itFn(test.l)
		if it.Insert(test.elems...); !Equal(test.l, test.want) {
			t.Errorf("Insert got %v, want %v", test.l, test.want)
		}
	}
}

var frwIterMutDeleteTests = []struct {
	l       *List[int]
	itFn    func(*List[int]) *FrwIterMut[int]
	hasNext bool
	next    int
	want    *List[int]
}{
	{
		l: New(1, 2, 7, 3),
		itFn: func(l *List[int]) *FrwIterMut[int] {
			it := IterMut(l)
			iters.Advance[*int](it, 3)
			return it
		},
		hasNext: true,
		next:    3,
		want:    New(1, 2, 3),
	},
}

func TestFrwIterMutDelete(t *testing.T) {
	for _, test := range frwIterMutDeleteTests {
		it := test.itFn(test.l)
		it.Delete()
		if !Equal(test.l, test.want) {
			t.Errorf("Delete got %v, want %v", test.l, test.want)
		}
		if it.HasNext() != test.hasNext {
			t.Errorf("it.HasNext() = %v, want = %v", it.HasNext(), test.hasNext)
		}
		if it.HasNext() {
			next := *it.Next()
			if next != test.next {
				t.Errorf("it.Next() = %v, want = %v", next, test.hasNext)
			}
		}
	}
}

var revIterMutDeleteTests = []struct {
	l       *List[int]
	itFn    func(*List[int]) *RevIterMut[int]
	hasNext bool
	next    int
	want    *List[int]
}{
	{
		l: New(1, 2, 3, 4),
		itFn: func(l *List[int]) *RevIterMut[int] {
			it := RIterMut(l)
			iters.Advance[*int](it, 1)
			return it
		},
		hasNext: true,
		next:    3,
		want:    New(1, 2, 3),
	},
}

func TestRevIterMutDelete(t *testing.T) {
	for _, test := range revIterMutDeleteTests {
		it := test.itFn(test.l)
		it.Delete()
		if !Equal(test.l, test.want) {
			t.Errorf("Delete got %v, want %v", test.l, test.want)
		}
		if it.HasNext() != test.hasNext {
			t.Errorf("it.HasNext() = %v, want = %v", it.HasNext(), test.hasNext)
		}
		if it.HasNext() {
			next := *it.Next()
			if next != test.next {
				t.Errorf("it.Next() = %v, want = %v", next, test.hasNext)
			}
		}
	}
}
