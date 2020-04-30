package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	var list []int

	for i := 0; i < 10; i++ {
		v := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutex.Lock()
			list = append(list, v)
			mutex.Unlock()
		}()
	}

	wg.Wait()
	fmt.Printf("%v\n%#v\n", len(list), list)
}
