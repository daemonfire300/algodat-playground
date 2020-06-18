package sorting

import (
	"math/rand"
)

func NotReallyQuicksort(lst []int) []int {
	if len(lst) <= 1 {
		return lst
	}
	pivotIdx := rand.Intn(len(lst))
	pivot := lst[pivotIdx]
	left := make([]int, 0)
	right := make([]int, 0)
	for i, v := range lst {
		if i == pivotIdx {
			continue
		}
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	left = append(left, pivot)
	res := make([]int, 0)
	res = append(res, NotReallyQuicksort(left)...)
	res = append(res, NotReallyQuicksort(right)...)
	return res
}

func Quicksort2(lst []int) []int {
	if len(lst) < 2 {
		return lst
	}

	left, right := 0, len(lst)-1

	pivot := rand.Int() % len(lst)
	lst[pivot], lst[right] = lst[right], lst[pivot]
	for i, _ := range lst {
		if lst[i] < lst[right] {
			lst[left], lst[i] = lst[i], lst[left]
			left++
		}
	}

	lst[left], lst[right] = lst[right], lst[left]
	Quicksort2(lst[:left])
	Quicksort2(lst[left+1:])
	return lst
}
func qsort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	// Pick a pivot
	pivotIndex := rand.Int() % len(a)

	// Move the pivot to the right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]

	// Go down the rabbit hole
	qsort(a[:left])
	qsort(a[left+1:])

	return a
}
