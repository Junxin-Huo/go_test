package main

import (
	"fmt"
	"sync"
)

type LRUCache struct {
	sync.RWMutex
	count int
	max   int
	data  map[int]string
	score []int
}

func (lru *LRUCache) updateScore(l int) {
	lru.score = append(lru.score, l)
}

func (lru *LRUCache) deleteScore(l int) {
	//lru.score.
}

func (lru *LRUCache) Get(l int) (string, bool) {
	lru.RLocker()
	defer lru.RUnlock()

	v, ok := lru.data[l]
	if ok {
		lru.updateScore(l)
		return v, ok
	}
	return "", false
}

func (lru *LRUCache) Put(l int, s string) {
	lru.Lock()
	defer lru.Unlock()

	if lru.count >= lru.max {
		delete(lru.data, 1)
		lru.deleteScore(l)
		lru.data[l] = s
	} else {
		lru.data[l] = s
		lru.count++
	}

	lru.updateScore(l)
}

func main () {
	data := []int{1,2,3,1,2,5,1,2,1}
	check := map[int]bool{}

	for _, v := range data {
		//fmt.Println(k, v)
		num, ok := check[v]
		if ok {
			check[v] = !num
		} else {
			check[v] = false
		}
	}

	for k, v := range check {
		fmt.Println(k, v)

	}
}





