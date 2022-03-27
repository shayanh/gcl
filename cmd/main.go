package main

import (
	"encoding/json"
	"fmt"

	"github.com/shayanh/gcl/iters"
	"github.com/shayanh/gcl/lists"
)

func printList[T any](lst *lists.List[T]) {
	fmt.Print("lists = ")
	iters.ForEach[T](lists.Iter(lst), func(t T) {
		fmt.Print(t, ", ")
	})
	fmt.Println()
}

func main() {
	lst := lists.New(1, 2, 7)
	lists.PushFront(lst, -9)
	printList(lst)

	reducer := iters.Reduce[int]
	sum := reducer(lists.Iter(lst), func(a, b int) int {
		return a + b
	})
	fmt.Println("sum =", sum)

	lists.Insert(lists.Iter(lst), 11, 12)
	printList(lst)

	lists.Insert(lists.RIter(lst), 13, 14)
	printList(lst)

	it := iters.Map[int, int](lists.Iter(lst), func(t int) int {
		return t * 2
	})
	lst2 := lists.FromIter(it)
	printList(lst2)

	f := func(t int) bool {
		return t%2 == 0
	}
	g := func(t int) {
		fmt.Print(t, ", ")
	}
	iters.ForEach(iters.Filter[int](lists.Iter(lst), f), g)
	fmt.Println()

	lists.Sort(lst)
	printList(lst)

	fmt.Println(lst)

	a, err := json.Marshal(lst)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(a))

	var lst3 *lists.List[int]
	err = json.Unmarshal(a, &lst3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(lst3)
}
