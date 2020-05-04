package main

import "fmt"

func f1() {
	a1 := []int{1, 2, 3, 4, 5, 6}
	a2 := a1
	a3 := a1[1:3]

	a1[1] = 999

	fmt.Println("a1=", a1, "a2=", a2, "a3=", a3)
}

func f2() {
	a1 := []int{1, 2, 3, 4, 5, 6}
	a2 := a1
	a3 := a1[1:3]

	a2 = append(a2, 888)

	a1[1] = 999

	fmt.Println("a1=", a1, "a2=", a2, "a3=", a3)
}

func f3() {
	a1 := []int{1, 2, 3, 4, 5, 6}
	a2 := a1
	a3 := a1[1:3]
	a2 = append(a2, 888)
	a3 = append(a3, 777)

	a1[1] = 999

	fmt.Println("a1=", a1, "a2=", a2, "a3=", a3)
}

func f4() {
	// 切片实质上是对底层匿名数组的引用
	slice := make([]int, 5, 5)
	slice1 := slice
	slice2 := slice[:]
	slice3 := slice[0:4]
	slice4 := slice[1:5]
	slice[1] = 1
	fmt.Println(slice)  //[0 1 0 0 0]
	fmt.Println(slice1) //[0 1 0 0 0]
	fmt.Println(slice2) //[0 1 0 0 0]
	fmt.Println(slice3) //[0 1 0 0]
	fmt.Println(slice4) //[1 0 0 0]
}

func f5() {
	// 当元素数量超过容量
	// 切片会在底层申请新的数组
	slice := make([]int, 5, 5)
	slice1 := slice
	slice = append(slice, 1)
	slice[0] = 1
	fmt.Println(slice)  //[1 0 0 0 0 1]
	fmt.Println(slice1) //[0 0 0 0 0]
	// copy 函数提供深拷贝功能
	// 但需要在拷贝前申请空间
	slice2 := make([]int, 4, 4)
	slice3 := make([]int, 5, 5)
	fmt.Println(copy(slice2, slice)) //4
	fmt.Println(copy(slice3, slice)) //5
	slice2[1] = 2
	slice3[1] = 3
	fmt.Println(slice)  //[1 0 0 0 0 1]
	fmt.Println(slice2) //[1 2 0 0]
	fmt.Println(slice3) //[1 3 0 0 0]
}

func main() {
	f1()
	f2()
	f3()
	f4()
	fmt.Println("f5():")
	f5()
}
