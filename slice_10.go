package main

import "fmt"

func func1() {
	s := []int{1, 2, 3, 4}
	m := make(map[int]*int)

	for k, v := range s {
		m[k] = &v
	}

	for key, value := range m {
		fmt.Printf("map[%v]=%v\n", key, *value)
	}
	fmt.Println(m)
}

func func2() {
	s := []int{1, 2, 3, 4}
	m := make(map[int]*int)

	for k, v := range s {
		n := v
		m[k] = &n
	}

	for key, value := range m {
		fmt.Printf("map[%v]=%v\n", key, *value)
	}
	fmt.Println(m)
}

func main() {
	fmt.Println("func1():")
	func1()
	fmt.Println("func2():")
	func2()
}
