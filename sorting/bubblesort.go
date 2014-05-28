package main

import (
	"fmt"
)

func main() {
	unsorted := getArrayFromInput()
	fmt.Println("Unsorted:", unsorted)
	sorted := bubbleSort(unsorted)
	fmt.Println("sorted:", sorted)
}

func getArrayFromInput() []int {
	var n int
	array := []int{}
	for {
		_, err := fmt.Scanf("%d", &n)
		if err == nil {
			array = append(array, n)
		} else {
			return array
		}
	}
}

func bubbleSort(unsorted []int) []int {
	a := make([]int, len(unsorted), cap(unsorted))
	copy(a, unsorted)
	sorted := false
	var buffer int
	for !sorted {
		sorted = true
		for i := 0; i < len(a)-1; i++ {
			if a[i] > a[i+1] {
				sorted = false
				buffer = a[i]
				a[i] = a[i+1]
				a[i+1] = buffer
			}
		}
	}
	return a
}
