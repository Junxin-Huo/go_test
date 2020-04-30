package main

import (
	"fmt"
)

func main() {

	ch := make(chan struct{})
	count := 2 // count 表示活动的协程个数

	go func() {
		fmt.Println("Goroutine 1")
		ch <- struct{}{} // 协程结束，发出信号
	}()

	go func() {
		fmt.Println("Goroutine 2")
		ch <- struct{}{} // 协程结束，发出信号
	}()

	for range ch {
		// 每次从ch中接收数据，表明一个活动的协程结束
		count--
		// 当所有活动的协程都结束时，关闭管道
		if count == 0 {
			close(ch)
		}
	}
}