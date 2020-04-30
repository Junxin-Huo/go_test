package main

import "fmt"

func main() {
	/* 创建切片 */
	numbers := []int{0,1,2,3,4,5,6,7,8}
	printSlice(numbers)

	numbers = make([]int,2,4)
	printSlice(numbers)
}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
