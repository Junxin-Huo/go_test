package main

import "fmt"

func main () {
	c := make(chan int)
	go func() { c <- 3 }()
	n := <-c
	defer fmt.Println(n)
	fmt.Println("abc")
}
