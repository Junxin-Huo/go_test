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

	a := p.Get().(int)
	fmt.Println(a)
	p.Put(1)
	p.Put(4)
	p.Put(2)
	p.Put(5)

	b := p.Get().(int)
	//runtime.GC()  // 执行回收操作后 1 0 0 0
	//time.Sleep(time.Second)
	c := p.Get().(int)
	d := p.Get().(int)
	fmt.Println(b, c, d, p.Get())
}
