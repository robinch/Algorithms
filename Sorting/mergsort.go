package main

import (
	"fmt"
)

func main() {
	divide([]int{1, 6, 3, 7})
}

func divide(list []int) []int {
	fmt.Println(list)
	if len(list) <= 1 {
		return list
	}
	n := len(list) / 2
	a := divide(list[0:n])
	b := divide(list[n:len(list)])
	merge(a, b)
	return nil
}

func merge(a, b []int) {
	x := len(a) + len(b)
	fmt.Println("X: ", x)
	var c []int
	c = append(a,b)
	fmt.Println("c: ", len(c))

	// for len(a) > 0 || len(b) > 0 {
	// 	fmt.Println("worked")
	// }
}
