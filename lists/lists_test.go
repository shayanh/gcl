package lists

import (
	"testing"
)

var reverseTests = []struct {
	l    *List[int]
	want *List[int]
}{
	{
		New[int](),
		New[int](),
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
			t.Errorf("Reverse(%v) = %v, want %v", test.l, cloned, test.want)
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
			t.Errorf("Compact(%v) = %v, want %v", test.l, cloned, test.want)
		}
	}
}
