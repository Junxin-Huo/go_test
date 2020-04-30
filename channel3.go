package main

import (
	"fmt"
	"time"
)

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s, %d", msg, i)
			time.Sleep(time.Second)
		}
	}()
	return c
}

func fanIn(input1, input2 <- chan string) <-chan string {
	c := make(chan string)
	go func() { for { c <- <- input1 }}()
	go func() { for { c <- <- input2 }}()
	return c
}

func main()  {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10 ; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring. I'm leaving")
}
