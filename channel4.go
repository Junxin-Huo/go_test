package main

import (
	"fmt"
	"time"
)

func main () {
	c := make(chan int, 1000)

	for i := 1; i <= 100; i++{
		c <- i
	}

	go func() {
		for {
			select {
			case a := <-c:
				fmt.Println(a)
			}
		}
	}()

	time.Sleep(time.Second)
	//len()
}
