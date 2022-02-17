package main

import (
	"fmt"

	"github.com/shayanh/gogl/iters"
	"github.com/shayanh/gogl/lists"
)

func printList[T any](lst *lists.List[T]) {
	fmt.Print("lists = ")
	iters.ForEach[T](lists.Iter(lst), func(t T) {
		fmt.Print(t, ", ")
	})
	fmt.Println()
}

func main() {
	lst := lists.NewList[int](1, 2, 7)
	lists.PushFront(lst, -9)
	printList(lst)

	sum := iters.Reduce[int](lists.Iter(lst), func(a, b int) int {
		return a + b
	})
	fmt.Println("sum =", sum)

	lists.Insert(lists.Iter(lst), 11, 12)
	printList(lst)

	lists.Insert(lists.RIter(lst), 13, 14)
	printList(lst)

	lst2 := lists.FromIter[int](iters.Map[int, int](lists.Iter(lst), func(t int) int {
		return t * 2
	}))
	printList(lst2)

	iters.ForEach[int](iters.Filter[int](lists.Iter(lst), 
		func (t int) bool {
			return t % 2 == 0
		}),
		func (t int) {
			fmt.Print(t, ", ")
		},
	)
	fmt.Println()
}
