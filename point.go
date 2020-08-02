package main

import "fmt"

func main() {
	p:=Person{"张三"}
	fmt.Printf("Before change p address: %p | %v\n", &p, p)
	modify(p)
	fmt.Printf("After change p address: %p | %v\n", &p, p)
}

type Person struct {
	Name string
}

func modify(p Person) {
	p.Name = "李四"
	fmt.Printf("Inside change p address: %p | %v\n", &p, p)
}
