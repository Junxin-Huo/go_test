package main

import "fmt"

func main() {
	s := make([]int, 0, 2)
	s = append(s, 0)
	s = append(s, 1)
	f(s)
	fmt.Println(s)
	f1(s)
	fmt.Println(s)
}

func f(s []int) {
	s[0] = 5
	s = append(s, 3)
	s = append(s, 4)
	s[0] = 7
	fmt.Println(s)
}

func f1(s []int) {
	s[0] = 6
	fmt.Println(s)
}
