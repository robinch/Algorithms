package main

import (
	"fmt"
)

func main() {
	list := []int{42, 25, 1, 2, 5, 77, 2, 3, 552}
	fmt.Println(mergeSort(list))
}

func mergeSort(unsortedList []int) []int {
	// Have to make a copy of the list so it does not
	// change the original
	list := make([]int, len(unsortedList))
	copy(list, unsortedList)
	if len(list) <= 1 {
		return list
	} else {
		n := len(list) / 2
		return merge(mergeSort(list[:n]), mergeSort(list[n:]))
	}

}

// Changed the function to a recursive one
// Much nicer code!
func merge(a, b []int) []int {
	if len(a) == 0 {
		return append(a, b...)
	} else if len(b) == 0 {
		return append(b, a...)
	} else if a[0] < b[0] {
		return append(a[:1], merge(a[1:], b)...)
	} else {
		return append(b[:1], merge(b[1:], a)...)
	}
}
