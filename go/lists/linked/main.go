package main

import "fmt"

type Item[T any] struct {
	val  T
	next *Item[T]
}

func print[T any](i *Item[T]) {
	fmt.Println(i.val)
	fmt.Println("|")
	if i.next == nil {
		fmt.Println(".")
	} else {
		print(i.next)
	}
}

func appendItem[T any](head *Item[T], val T) {
	if head.next == nil {
		head.next = &Item[T]{
			val: val,
		}
		return
	}
	appendItem(head.next, val)
}

func main() {
	lst := &Item[int]{
		val:  1,
		next: nil,
	}
	print(lst)
	fmt.Println("---")
	appendItem(lst, 2)
	appendItem(lst, 3)
	appendItem(lst, 4)
	appendItem(lst, 5)
	appendItem(lst, 6)
	print(lst)
}
