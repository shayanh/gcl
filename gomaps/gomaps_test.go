package gomaps_test

import (
	"testing"

	"github.com/shayanh/gcl"
	"github.com/shayanh/gcl/gomaps"
	"github.com/shayanh/gcl/goslices"
	"github.com/shayanh/gcl/iters"
)

func TestIter(t *testing.T) {
	m := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
	}
	it := gomaps.Iter(m)
	want := goslices.Iter([]gcl.MapElem[string, int]{
		{Key: "1", Value: 1},
		{Key: "2", Value: 2},
		{Key: "3", Value: 3},
	})
	if !iters.Equal[gcl.MapElem[string, int]](it, want) {
		t.Error("Wrong Iter result")
	}
}
