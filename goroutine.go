package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("main defer")
	go func() {
		defer fmt.Println("goroutine defer")
		for {
			time.Sleep(1 * time.Second)
			fmt.Println("goroutine")
		}
	}()
	time.Sleep(10 * time.Second)
	fmt.Println("main end")
}
