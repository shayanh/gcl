package main

import (
	"fmt"

	"github.com/shayanh/gogl/list"
)

func main() {
	lst := list.NewList[int]()
	lst.PushBack(1)
	lst.PushBack(2)
	lst.PushBack(7)
	lst.PushFront(-9)

	fmt.Println("foreach")
	lst.ForEach(func(t int) {
		fmt.Println(t)
	})

	lst.Reverse()
	fmt.Println("after reverse")
	lst.ForEach(func (t int) {
		fmt.Println(t)
	})

	it := lst.Iter()
	it.Next()
	lst.Insert(it, 100)

	fmt.Println("after insert")
	lst.ForEach(func (t int) {
		fmt.Println(t)
	})

	it = lst.Iter()
	it.Next()
	fmt.Println("it:", it.Value())

	it2 := lst.Erase(it)
	fmt.Println("it2:", it2.Value())

	fmt.Println("after erase")
	lst.ForEach(func (t int) {
		fmt.Println(t)
	})

	fmt.Println("manual iteration with for loop")
	for it := lst.Iter(); !it.Done(); it.Next() {
		fmt.Println(it.Value())
	}

	fmt.Println("reverse iteration with for loop")
	for it := lst.RIter(); !it.Done(); it.Next() {
		fmt.Println(it.Value())
	}
}
