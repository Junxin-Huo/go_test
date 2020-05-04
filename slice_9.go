package main

import "fmt"

// 数组是 值类型 。函数传递会拷贝，不会更改原值
// 切片是 引用类型 。函数传递会传指针，会更改原值
func main() {
	numbers := []int{1, 2, 3, 4, 5}
	printSlice(numbers)
	addOne(numbers)
	printSlice(numbers)

	arrays := [5]int{1, 2, 3, 4, 5}
	printArray(arrays)
	addOne2(arrays)
	printArray(arrays)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func addOne(x []int) {
	for i := range x {
		x[i] = x[i] + 1
	}
}

func printArray(x [5]int) {
	fmt.Println(x)
}

func addOne2(x [5]int) {
	for i := range x {
		x[i] = x[1] + 1
	}
}
