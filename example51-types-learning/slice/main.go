package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("len: ", len(slice))
	fmt.Println("cap: ", cap(slice))

	//改变切片长度
	slice = append(slice, 6)
	fmt.Println("after append operation: ")
	fmt.Println("len: ", len(slice))
	fmt.Println("cap: ", cap(slice)) //注意，底层数组容量不够时，会重新分配数组空间，通常为两倍
}
