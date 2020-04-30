package main

import (
	"fmt"
	"sync"
)

func main() {
	var list []int
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			list = append(list, i)
		}(i)
	}
	wg.Wait()
	fmt.Printf("%v\n%#v\n", len(list), list)
}
