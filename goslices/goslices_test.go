package goslices

import (
	"testing"

	"golang.org/x/exp/slices"
)

var reverseTests = []struct {
	s    []int
	want []int
}{
	{
		nil,
		nil,
	},
	{
		[]int{1, 2, 3},
		[]int{3, 2, 1},
	},
	{
		[]int{1, 2, 3, 4},
		[]int{4, 3, 2, 1},
	},
	{
		[]int{1, 2, 1, 3},
		[]int{3, 1, 2, 1},
	},
}

func TestReverse(t *testing.T) {
	for _, test := range reverseTests {
		cloned := slices.Clone(test.s)
		if Reverse(cloned); !slices.Equal(cloned, test.want) {
			t.Errorf("Reverse(%v) got %v, want %v", test.s, cloned, test.want)
		}
	}
}
