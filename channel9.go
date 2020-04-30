package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	ch := make(chan int, 10)

	// 4个协程消费数据
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for t := range ch {
				fmt.Println(t)
			}
		}()
	}

	// 插入数据
	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	// 等待执行完毕
	wg.Wait()
}
