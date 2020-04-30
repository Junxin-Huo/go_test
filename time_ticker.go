package main

import (
	"fmt"
	"time"
)

func LoopTask(t time.Time) {
	go func() {
		for {
			time.Sleep(time.Duration(5) * time.Second)
			fmt.Println(t, "end")
		}
	}()
}
func main() {

	go func() {
		defer fmt.Println("in go func")
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		count := 1
		for {
			fmt.Println("go func!")
			time.Sleep(time.Duration(1) * time.Second)
			if count == 10 {
				time.Sleep(time.Duration(10) * time.Second)
			}
			if count == 11 {
				panic("test go func panic")
			}
			count++
		}
	}()

	ticker := time.NewTicker(time.Second * time.Duration(2))
	for {
		select {
		case t := <-ticker.C:
			fmt.Println(t, "begin")
			LoopTask(t)
		}
	}
}
