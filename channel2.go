package main

import (
	"fmt"
	"time"
)

type Ball struct{hits int}

func main() {
	table := make(chan *Ball)
	go player("ping", table)
	go player("pang", table)

	table <- new(Ball)
	time.Sleep(1 * time.Second)

	ball := <-table
	fmt.Println("grab")

	time.Sleep(1 * time.Second)

	table <- ball
	fmt.Println("emission")

	time.Sleep(1 * time.Second)
	<-table

	//panic("show me the stacks")
}

func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
