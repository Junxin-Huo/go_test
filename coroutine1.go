package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	runtime.GOMAXPROCS(1)
	go func() {
		wg.Add(1)
		for i := 0; i < 5; i++ {
			fmt.Printf("--- 1 ---, %v, %v\n", i, time.Now())
			time.Sleep(time.Second)
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		for i := 0; i < 8; i++ {
			fmt.Printf("--- 2 ---, %v, %v\n", i, time.Now())
			time.Sleep(time.Second)
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		for i := 0; i < 6; i++ {
			fmt.Printf("--- 3 ---, %v, %v\n", i, time.Now())
			time.Sleep(time.Second)
		}
		wg.Done()
	}()

	time.Sleep(time.Second)
	wg.Wait()
}
