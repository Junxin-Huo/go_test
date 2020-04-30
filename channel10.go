package main

import (
	"fmt"
)

func main() {

	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3

	//关闭通道后，可以继续读
	close(ch)
	for value := range ch {
		fmt.Println("value:", value)
	}

}
