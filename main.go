package main

import "fmt"

func main() {
	l := NewList[int]()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(7)

	fmt.Println("foreach")
	l.ForEach(func (t int) {
		fmt.Println(t)
	})

	fmt.Println("foreach rev iter")
	l.ForEachIter(l.RevIter(), func (t int) {
		fmt.Println(t)
	})

	l.Reverse()
	fmt.Println("after reverse")
	l.ForEach(func (t int) {
		fmt.Println(t)
	})
}
