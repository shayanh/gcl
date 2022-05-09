package iters_test

import (
	"strconv"
	"testing"

	"github.com/shayanh/gcl"
	"github.com/shayanh/gcl/goslices"
	"github.com/shayanh/gcl/iters"
)

var equalTests = []struct {
	it1, it2 func() iters.Iterator[int]
	want     bool
}{
	{
		func() iters.Iterator[int] { return goslices.Iter[[]int](nil) },
		func() iters.Iterator[int] { return goslices.Iter[[]int](nil) },
		true,
	},
	{
		func() iters.Iterator[int] { return goslices.Iter([]int{1, 2, 3}) },
		func() iters.Iterator[int] { return goslices.Iter([]int{1, 2, 3}) },
		true,
	},
	{
		func() iters.Iterator[int] { return goslices.Iter([]int{1, 1, 1}) },
		func() iters.Iterator[int] { return goslices.Iter([]int{1, 2, 1}) },
		false,
	},
	{
		func() iters.Iterator[int] { return goslices.Iter([]int{1, 2, 3}) },
		func() iters.Iterator[int] { return goslices.Iter([]int{1, 2, 3, 4}) },
		false,
	},
}

func TestEqual(t *testing.T) {
	for _, test := range equalTests {
		if res := iters.Equal(test.it1(), test.it2()); res != test.want {
			t.Errorf("Equal(it1, it2) = %v, want = %v", res, test.want)
		}
	}
}

func TestEqualFunc(t *testing.T) {
	for _, test := range equalTests {
		if res := iters.EqualFunc(test.it1(), test.it2(), gcl.Equal[int]); res != test.want {
			t.Errorf("Equal(it1, it2) = %v, want = %v", res, test.want)
		}
	}
}

func TestMap(t *testing.T) {
	arr := []int{1, 2, 3}
	func() {
		it := goslices.Iter(arr)
		got := iters.Map[int](it, func(t int) int {
			return t * 2
		})
		want := goslices.Iter([]int{2, 4, 6})
		if !iters.Equal[int](got, want) {
			t.Errorf("Wrong Map result")
		}
		if it.HasNext() {
			t.Errorf("it.HasNext() must be false")
		}
	}()
	func() {
		it := goslices.Iter(arr)
		got := iters.Map[int](it, func(t int) string {
			return strconv.Itoa(t)
		})
		want := goslices.Iter([]string{"1", "2", "3"})
		if !iters.Equal[string](got, want) {
			t.Errorf("Wrong Map result")
		}
		if it.HasNext() {
			t.Errorf("it.HasNext() must be false")
		}
	}()
}

var reduceTests = []struct {
	it   iters.Iterator[int]
	fn   func(int, int) int
	want int
}{
	{
		goslices.Iter([]int{1, 2, 3}),
		func(a, b int) int {
			return a + b
		},
		6,
	},
	{
		goslices.Iter([]int{1, 2, 3, 4}),
		func(a, b int) int {
			return a * b
		},
		24,
	},
}

func TestReduce(t *testing.T) {
	for _, test := range reduceTests {
		got := iters.Reduce(test.it, test.fn)
		if got != test.want {
			t.Errorf("Wrong Reduce result")
		}
		if test.it.HasNext() {
			t.Errorf("it.HasNext() must be false")
		}
	}
}

var foldIntTests = []struct {
	it   iters.Iterator[int]
	fn   func(int, int) int
	init int
	want int
}{
	{
		goslices.Iter([]int{1, 2, 3}),
		func(a, b int) int {
			return a + b
		},
		0,
		6,
	},
	{
		goslices.Iter([]int{1, 2, 3, 4}),
		func(a, b int) int {
			return a * b
		},
		1,
		24,
	},
}

var foldStrTests = []struct {
	it   iters.Iterator[int]
	fn   func(string, int) string
	init string
	want string
}{
	{
		goslices.Iter([]int{1, 2, 3}),
		func(acc string, a int) string {
			return acc + strconv.Itoa(a)
		},
		"",
		"123",
	},
	{
		goslices.RIter([]int{1, 2, 3}),
		func(acc string, a int) string {
			return acc + strconv.Itoa(a)
		},
		"",
		"321",
	},
}

func TestFold(t *testing.T) {
	for _, test := range foldIntTests {
		got := iters.Fold(test.it, test.fn, test.init)
		if got != test.want {
			t.Errorf("Wrong Fold result")
		}
		if test.it.HasNext() {
			t.Errorf("it.HasNext() must be false")
		}
	}
	for _, test := range foldStrTests {
		got := iters.Fold(test.it, test.fn, test.init)
		if got != test.want {
			t.Errorf("Wrong Fold result")
		}
		if test.it.HasNext() {
			t.Errorf("it.HasNext() must be false")
		}
	}
}

var filterTests = []struct {
	it   iters.Iterator[int]
	fn   func(int) bool
	want iters.Iterator[int]
}{
	{
		goslices.Iter([]int{1, 2, 3, 4}),
		func(n int) bool {
			return n%2 == 0
		},
		goslices.Iter([]int{2, 4}),
	},
	{
		goslices.Iter([]int{}),
		func(n int) bool {
			return n%2 == 0
		},
		goslices.Iter([]int{}),
	},
}

func TestFilter(t *testing.T) {
	for _, test := range filterTests {
		got := iters.Filter(test.it, test.fn)
		if !iters.Equal(got, test.want) {
			t.Errorf("Wrong Filter result")
		}
		if test.it.HasNext() {
			t.Errorf("it.HasNext() must be false")
		}
	}
}

var findTests = []struct {
	it      iters.Iterator[int]
	fn      func(int) bool
	wantVal int
	wantOk  bool
}{
	{
		goslices.Iter([]int{1, 2, 3, 4}),
		func(n int) bool {
			return n%2 == 0
		},
		2,
		true,
	},
	{
		goslices.Iter([]int{}),
		func(n int) bool {
			return n%2 == 0
		},
		0,
		false,
	},
}

func TestFind(t *testing.T) {
	for _, test := range findTests {
		val, ok := iters.Find(test.it, test.fn)
		if val != test.wantVal || ok != test.wantOk {
			t.Errorf("Wrong Find result")
		}
	}
}

func TestZip(t *testing.T) {
	it1 := goslices.Iter([]int{1, 2})
	it2 := goslices.Iter([]string{"a", "b", "c"})
	zipped := iters.Zip[int, string](it1, it2)

	want := goslices.Iter([]gcl.Zipped[int, string]{
		{First: 1, Second: "a"},
		{First: 2, Second: "b"},
	})

	if !iters.Equal[gcl.Zipped[int, string]](zipped, want) {
		t.Error("Wrong Zip result")
	}
	if it1.HasNext() {
		t.Errorf("it1.HasNext() must be false")
	}
	if !it2.HasNext() {
		t.Errorf("it2.HasNext() must be true")
	}
}
