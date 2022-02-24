package lists

import (
	"fmt"
	"testing"
)

func TestCompact(t *testing.T) {
	lst := New(1, 1, 1, 2)
	Compact(lst)
	fmt.Println(lst)
}
