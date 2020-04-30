package main

import (
	"fmt"
	"sync"
)

func main() {
	println("start main")
	wg := sync.WaitGroup{}
	ch := make(chan int, 4)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		// 如果不关闭channel,会引发panic
		close(ch)
	}()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			for v := range ch {
				fmt.Println(i, v)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	println("end main")
}
