package main

import (
	"fmt"
	"strconv"

	"github.com/shayanh/gogl/list"
	"github.com/shayanh/gogl/iters"
)

func printList[T any](lst *list.List[T]) {
	fmt.Print("list = ")
	iters.ForEach[T](list.Begin(lst), func(t T) {
		fmt.Print(t, ", ")
	})
	fmt.Println()
}

func main() {
	lst := list.NewList[int](1, 2, 7)
	list.PushFront(lst, -9)
	printList(lst)

	list.Insert(lst, list.Begin(lst), 11, 12)
	printList(lst)

	list.Insert(lst, list.RBegin(lst), 13, 14)
	printList(lst)

	lst2 := list.NewList[int](iters.Map[int, int](list.Begin(lst), func(t int) int {
		return t * 2
	})...)
	printList(lst2)

	lst3 := list.NewList[string](iters.Map[int, string](list.Begin(lst), func(t int) string {
		return strconv.Itoa(t) + "a"
	})...)
	printList(lst3)
}
