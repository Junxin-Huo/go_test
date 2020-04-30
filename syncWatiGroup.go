package main

import (
	"fmt"
	"sync"
	"time"
)

func handle(data int) {
	fmt.Println(data)
	time.Sleep(time.Second * 6)
	fmt.Println("in handle")
}


func main() {
	var wg sync.WaitGroup
	wg.Add(16)
	for i := 0; i < 16; i++ {
		go func(data int) {
			handle(data)
			wg.Done()
		}(i)
	}
	fmt.Println("now wait.....")
	wg.Wait()
	fmt.Println("end")
}