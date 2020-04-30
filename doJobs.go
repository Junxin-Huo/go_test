package main

import (
	"fmt"
	"sync"
)

type work struct {
	wg     sync.WaitGroup
	jobs   chan *string
	gCount int
}

func (w *work) Init() {
}

func (w *work) Run(task []string) {
	w.wg.Add(1)
	go func() {
		// range的v，只是拷贝，直接用会有问题，建议通过k索引访问数据
		//for _, v := range task {
		//	tmp := v
		//	w.jobs <- &tmp
		//}
		for k, _ := range task {
			w.jobs <- &task[k]
		}
		close(w.jobs)
		w.wg.Done()
	}()
	//并发数
	for i := 0; i <= w.gCount; i++ {
		w.wg.Add(1)
		go func() {
			defer w.wg.Done()
			for t := range w.jobs {
				fmt.Println(*t)
			}
		}()
	}
}

func (w *work) Wait() {
	w.wg.Wait()
}

func main() {
	w := work{
		wg:     sync.WaitGroup{},
		jobs:   make(chan *string),
		gCount: 2,
	}
	ts := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	w.Run(ts)
	w.Wait()
	//time.Sleep(1)
}
