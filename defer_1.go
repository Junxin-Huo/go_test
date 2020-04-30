package main

import (
	"fmt"
	"strconv"
)

func a() {
	i := 0
	fmt.Println("a:" + strconv.Itoa(i))
	defer fmt.Println("b:" + strconv.Itoa(i))
	i++
	fmt.Println("c:" + strconv.Itoa(i))
	defer fmt.Println("d:" + strconv.Itoa(i))
	return
}

func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

func c() (i int) {
	defer func() { i++ }()
	return 1
}

func main() {
	a()

	b()

	fmt.Print("\nC:")
	fmt.Println(c())
}
