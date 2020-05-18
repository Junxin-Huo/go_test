package main

import (
	"fmt"
	"sync"
)

func main() {
	p := &sync.Pool{
		New: func() interface{} {
			return 0
		},
	}

	p.Put(1)
	fmt.Println(p.Get())
	fmt.Println(p.Get())


	p.Put(2)
	p.Put(3)
	p.Put(4)
	p.Put(5)

	fmt.Println(p.Get())
	fmt.Println(p.Get())
	fmt.Println(p.Get())
	fmt.Println(p.Get())
	fmt.Println(p.Get())
	fmt.Println(p.Get())
}
