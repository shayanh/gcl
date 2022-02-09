package main

import (
	"fmt"
	"strconv"

	"github.com/shayanh/gogl/list"
	"github.com/shayanh/gogl/iters"
)

func printList[T any](lst *list.List[T]) {
	fmt.Print("list = ")
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

	lst2 := list.NewList[int]()
	iters.Map[int, int](lst.Iter(), func(t int) int {
		return t * 2
	}, lst2.PushBack)
	printList(lst2)

	lst3 := list.NewList[string]()
	iters.Map[int, string](lst.Iter(), func(t int) string {
		return strconv.Itoa(t) + "a"
	}, lst3.PushBack)
	printList(lst3)
}
