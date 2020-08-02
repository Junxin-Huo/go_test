package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
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

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Printf("--- 4 ---, %v\n", time.Now())
		for true {}
	}()

	time.Sleep(time.Second)
	wg.Wait()
}
