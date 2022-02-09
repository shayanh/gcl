package main

import (
	"fmt"

	"github.com/shayanh/gogl/list"
)

func printList[T any](lst *list.List[T]) {
	fmt.Print("lst = ")
	lst.ForEach(func(t T) {
		fmt.Print(t, ", ")
	})
	fmt.Println()
}

func main() {
	lst := list.NewList[int](1, 2, 7)
	lst.PushFront(-9)
	printList(lst)

	lst.Insert(lst.Iter(), 11, 12)
	printList(lst)

	lst.Insert(lst.RIter(), 13, 14)
	printList(lst)
}
