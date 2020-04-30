package main

import "fmt"

func main() {
	a := new(int)
	//tmp := 0
	//a = &tmp
	//*a = 0;
	fmt.Println(a, *a)

	b := make(chan int)
	defer close(b)
	go func() { b <- 2 }()
	//b <- 1
	//<- b
	fmt.Println(b, <- b)
}
