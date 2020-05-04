package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6}
	s1 := arr[1:3]
	printSlice(s1)
	s2 := arr[1:3:4]
	printSlice(s2)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
