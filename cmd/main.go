package main

import (
	"fmt"
	"strconv"

	"github.com/shayanh/gogl/iters"
	"github.com/shayanh/gogl/lists"
)

func printList[T any](lst *lists.List[T]) {
	fmt.Print("lists = ")
	iters.ForEach[T](lists.Begin(lst), func(t T) {
		fmt.Print(t, ", ")
	})
	fmt.Println()
}

func main() {
	lst := lists.NewList[int](1, 2, 7)
	lists.PushFront(lst, -9)
	printList(lst)

	sum := iters.Reduce[int](lists.Begin(lst), func(a, b int) int {
		return a + b
	})
	fmt.Println("sum =", sum)

	lists.Insert(lst, lists.Begin(lst), 11, 12)
	printList(lst)

	lists.Insert(lst, lists.RBegin(lst), 13, 14)
	printList(lst)

	lst2 := lists.NewList[int](iters.Map[int, int](lists.Begin(lst), func(t int) int {
		return t * 2
	})...)
	printList(lst2)

	it := lists.Begin(lst)
	fmt.Println("it.Done() =", it.Done())
	lst3 := lists.NewList[string](iters.Map[int, string](it, func(t int) string {
		return strconv.Itoa(t) + "a"
	})...)
	printList(lst3)

	fmt.Println("it.Done() =", it.Done())
}
