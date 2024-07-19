package main

import "fmt"

/*
*
	change one ,affect two
*/
func main() {
	slice := []int{1, 2, 3, 4, 5}
	newSlice := slice[0:3]
	fmt.Println("before modifying underlying array:")
	fmt.Println("slice: ", slice)
	fmt.Println("newSlice: ", newSlice)
	fmt.Println()

	newSlice[0] = 6
	fmt.Println("after modifying underlying array:")
	fmt.Println("slice: ", slice)
	fmt.Println("newSlice: ", newSlice)
}
