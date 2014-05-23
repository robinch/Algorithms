package main

import (
	"fmt"
)

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{5, 6, 7}
	slice := append(slice1, slice2...)
	fmt.Println(slice)
	slice3 := make([]int, len(slice))
	copy(slice3, slice)
	slice[2] = 50000
	fmt.Println(slice3)
	fmt.Println(slice)
}
