package lists

import (
	"strings"
	"testing"

	"github.com/shayanh/gcl"
	"github.com/shayanh/gcl/iters"
)

var equalTests = []struct {
	l1, l2 *List[int]
	want   bool
}{
	{
		New[int](),
		New[int](),
		true,
	},
	{
		New(1, 2, 3),
		New(1, 2, 3),
		true,
	},
	{
		New(1, 1, 1),
		New(1, 2, 1),
		false,
	},
	{
		New(1, 2, 3),
		New(1, 2, 3, 4),
		false,
	},
}

func TestEqual(t *testing.T) {
	for _, test := range equalTests {
		if res := Equal(test.l1, test.l2); res != test.want {
			t.Errorf("Equal(%v, %v) = %v, want = %v", test.l1, test.l2, res, test.want)
		}
	}
}

func TestEqualFunc(t *testing.T) {
	for _, test := range equalTests {
		if res := EqualFunc(test.l1, test.l2, gcl.Equal[int]); res != test.want {
			t.Errorf("EqualFunc(%v, %v, gcl.Equal[int]) = %v, want = %v", test.l1, test.l2, res, test.want)
		}
	}

	l1 := New("A", "B", "C")
	l2 := New("a", "b", "c")
	if res := EqualFunc(l1, l2, strings.EqualFold); !res {
		t.Errorf("Equal(%v, %v, strings.EqualFold) = %v, want = %v", l1, l2, res, true)
	}
}

var compareTests = []struct {
	l1, l2 *List[int]
	want   int
}{
	{
		New(1, 2, 3),
		New(1, 2, 3, 4),
		-1,
	},
	{
		New(1, 2, 3, 4),
		New(1, 2, 3),
		+1,
	},
	{
		New(1, 2, 3),
		New(1, 4, 3),
		-1,
	},
	{
		New(1, 4, 3),
		New(1, 2, 3),
		+1,
	},
}

func TestCompare(t *testing.T) {
	for _, test := range equalTests {
		if res := Compare(test.l1, test.l2); (res == 0) != test.want {
			t.Errorf("Compare(%v, %v) = %v, want = %v", test.l1, test.l2, res, func(want bool) string {
				if want {
					return "0"
				}
				return "!= 0"
			}(test.want))
		}
	}

	for _, test := range compareTests {
		if res := Compare(test.l1, test.l2); res != test.want {
			t.Errorf("Compare(%v, %v) = %v, want = %v", test.l1, test.l2, res, test.want)
		}
	}
}

func TestCompareFunc(t *testing.T) {
	for _, test := range equalTests {
		if res := CompareFunc(test.l1, test.l2, gcl.Compare[int]); (res == 0) != test.want {
			t.Errorf("CompareFunc(%v, %v, gcl.Compare[int]) = %v, want = %v", test.l1, test.l2, res, func(want bool) string {
				if want {
					return "0"
				}
				return "!= 0"
			}(test.want))
		}
	}

	for _, test := range compareTests {
		if res := CompareFunc(test.l1, test.l2, gcl.Compare[int]); res != test.want {
			t.Errorf("CompareFunc(%v, %v, gcl.Compare[int]) = %v, want = %v", test.l1, test.l2, res, test.want)
		}
	}
}

var pushBackTests = []struct {
	l     *List[int]
	elems []int
	want  *List[int]
}{
	{
		New[int](),
		[]int{1, 2, 3},
		New(1, 2, 3),
	},
	{
		New(3, 4, 5),
		[]int{6},
		New(3, 4, 5, 6),
	},
}

func TestPushBack(t *testing.T) {
	for _, test := range pushBackTests {
		cloned := Clone(test.l)
		if PushBack(cloned, test.elems...); !Equal(cloned, test.want) {
			t.Errorf("PushBack(%v, %v) got %v, want %v", test.l, test.elems, cloned, test.want)
		}
	}
}

var pushFrontTests = []struct {
	l     *List[int]
	elems []int
	want  *List[int]
}{
	{
		New[int](),
		[]int{1, 2, 3},
		New(1, 2, 3),
	},
	{
		New(3, 4, 5),
		[]int{6, 7},
		New(6, 7, 3, 4, 5),
	},
}

func TestPushFront(t *testing.T) {
	for _, test := range pushFrontTests {
		cloned := Clone(test.l)
		if PushFront(cloned, test.elems...); !Equal(cloned, test.want) {
			t.Errorf("PushFront(%v, %v) got %v, want %v", test.l, test.elems, cloned, test.want)
		}
	}
}

var insertTests = []struct {
	l     *List[int]
	itFn  func(*List[int]) Iterator[int]
	elems []int
	want  *List[int]
}{
	{
		l: New(1, 3),
		itFn: func(l *List[int]) Iterator[int] {
			it := Iter(l)
			iters.Advance[int](it, 1)
			return it
		},
		elems: []int{2},
		want:  New(1, 2, 3),
	},
	{
		l: New(1, 2, 5, 6),
		itFn: func(l *List[int]) Iterator[int] {
			it := Iter(l)
			iters.Advance[int](it, 2)
			return it
		},
		elems: []int{3, 4},
		want:  New(1, 2, 3, 4, 5, 6),
	},
	{
		l: New(1, 2, 5, 6),
		itFn: func(l *List[int]) Iterator[int] {
			it := RIter(l)
			iters.Advance[int](it, 2)
			return it
		},
		elems: []int{3, 4},
		want:  New(1, 2, 3, 4, 5, 6),
	},
}

func TestInsert(t *testing.T) {
	for _, test := range insertTests {
		it := test.itFn(test.l)
		if Insert(it, test.elems...); !Equal(test.l, test.want) {
			t.Errorf("Insert got %v, want %v", test.l, test.want)
		}
	}
}

var deleteTests = []struct {
	l       *List[int]
	itFn    func(*List[int]) Iterator[int]
	hasNext bool
	next    int
	want    *List[int]
}{
	{
		l: New(1, 2, 7, 3),
		itFn: func(l *List[int]) Iterator[int] {
			it := Iter(l)
			iters.Advance[int](it, 3)
			return it
		},
		hasNext: true,
		next:    3,
		want:    New(1, 2, 3),
	},
	{
		l: New(1, 2, 3, 4),
		itFn: func(l *List[int]) Iterator[int] {
			it := RIter(l)
			iters.Advance[int](it, 1)
			return it
		},
		hasNext: true,
		next:    3,
		want:    New(1, 2, 3),
	},
}

func TestDelete(t *testing.T) {
	for _, test := range deleteTests {
		it := test.itFn(test.l)
		newIt := Delete(it)
		if !Equal(test.l, test.want) {
			t.Errorf("Delete got %v, want %v", test.l, test.want)
		}
		if !it.Valid() {
			t.Errorf("iterator should be invalidated after Delete")
		}
		if newIt.HasNext() != test.hasNext {
			t.Errorf("newIt.HasNext() = %v, want = %v", newIt.HasNext(), test.hasNext)
		}
		if newIt.HasNext() {
			next := newIt.Next()
			if next != test.next {
				t.Errorf("newIt.Next() = %v, want = %v", next, test.hasNext)
			}
		}
	}
}

var sortInts = New(38, 44, -90, -23, 14, -62, 34, 50, 25, 50)

func TestSort(t *testing.T) {
	cloned := Clone(sortInts)
	if Sort(cloned); !IsSorted(cloned) {
		t.Errorf("Sort(%v) got %v", sortInts, cloned)
	}
}

func TestSortFunc(t *testing.T) {
	cloned := Clone(sortInts)
	if SortFunc(cloned, gcl.Less[int]); !IsSortedFunc(cloned, gcl.Less[int]) {
		t.Errorf("SortFunc(%v, gcl.Less[int]) got %v", sortInts, cloned)
	}

	cloned = Clone(sortInts)
	if SortFunc(cloned, gcl.Greater[int]); !IsSortedFunc(cloned, gcl.Greater[int]) {
		t.Errorf("SortFunc(%v, gcl.Greater[int]) got %v", sortInts, cloned)
	}
}

var reverseTests = []struct {
	l    *List[int]
	want *List[int]
}{
	{
		New[int](),
		New[int](),
	},
	{
		New(1, 2, 3),
		New(3, 2, 1),
	},
	{
		New(1, 2, 3, 4),
		New(4, 3, 2, 1),
	},
	{
		New(1, 2, 1, 3),
		New(3, 1, 2, 1),
	},
}

func TestReverse(t *testing.T) {
	for _, test := range reverseTests {
		cloned := Clone(test.l)
		if Reverse(cloned); !Equal(cloned, test.want) {
			t.Errorf("Reverse(%v) got %v, want %v", test.l, cloned, test.want)
		}
	}
}

var compactTests = []struct {
	l    *List[int]
	want *List[int]
}{
	{
		New[int](),
		New[int](),
	},
	{
		New(1, 1, 1, 2),
		New(1, 2),
	},
	{
		New(1, 2, 3),
		New(1, 2, 3),
	},
	{
		New(1, 2, 2, 3, 3, 4),
		New(1, 2, 3, 4),
	},
}

func TestCompact(t *testing.T) {
	for _, test := range compactTests {
		cloned := Clone(test.l)
		if Compact(cloned); !Equal(cloned, test.want) {
			t.Errorf("Compact(%v) got %v, want %v", test.l, cloned, test.want)
		}
	}
}

func TestCompactFunc(t *testing.T) {
	for _, test := range compactTests {
		cloned := Clone(test.l)
		if CompactFunc(cloned, gcl.Equal[int]); !Equal(cloned, test.want) {
			t.Errorf("CompactFunc(%v, gcl.Equal[int]) got %v, want %v", test.l, cloned, test.want)
		}
	}
}

var indexTests = []struct {
	l    *List[int]
	v    int
	want int
}{
	{
		New(1, 2, 3, 4),
		3,
		2,
	},
	{
		New(1, 2, 10, 4),
		3,
		-1,
	},
}

func TestIndex(t *testing.T) {
	for _, test := range indexTests {
		res := Index(test.l, test.v)
		if res != test.want {
			t.Errorf("Index(%v, %v) = %v, want %v", test.l, test.v, res, test.want)
		}
	}
}

var indexFuncTests = []struct {
	l    *List[int]
	pred func(int) bool
	want int
}{
	{
		l: New(1, 2, 3, 4),
		pred: func(n int) bool {
			return n%2 == 0
		},
		want: 1,
	},
	{
		l: New(0, 2, 6, 4),
		pred: func(n int) bool {
			return n%2 == 1
		},
		want: -1,
	},
}

func TestIndexFunc(t *testing.T) {
	for _, test := range indexFuncTests {
		res := IndexFunc(test.l, test.pred)
		if res != test.want {
			t.Errorf("IndexFunc(%v, test.pred) = %v, want %v", test.l, res, test.want)
		}
	}
}

var posTests = []struct {
	l       *List[int]
	v       int
	found   bool
	hasNext bool
	next    int
}{
	{
		New(1, 2, 3, 4),
		3,
		true,
		true,
		4,
	},
	{
		New(1, 2, 10, 4),
		3,
		false,
		false,
		-1,
	},
}

func TestPos(t *testing.T) {
	for _, test := range posTests {
		it, found := Pos(test.l, test.v)
		if found != test.found {
			t.Errorf("Pos(%v, %v) found = %v, want %v", test.l, test.v, found, test.found)
		}
		if it.HasNext() != test.hasNext {
			t.Errorf("Pos(%v, %v) iterator HasNext = %v, want %v", test.l, test.v, it.HasNext(), test.hasNext)
		}
		if it.HasNext() {
			next := it.Next()
			if next != test.next {
				t.Errorf("Pos(%v, %v) iterator HasNext = %v, want %v", test.l, test.v, next, test.next)
			}
		}
	}
}
