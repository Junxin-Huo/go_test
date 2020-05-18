package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan interface{})
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go func(n int) {
			wg.Add(1)
			for {
				select {
				case <-ch:
					fmt.Printf("close %v\n", n)
					wg.Done()
					return
				}
			}
		}(i)
	}

	fmt.Println("begin")
	time.Sleep(3 * time.Second)
	fmt.Println("close")
	close(ch)
	fmt.Println("end")
	wg.Wait()
}
