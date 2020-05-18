package main

import "fmt"

func main() {
	var data []int
	for i := 0; i < 30; i++ {
		data = append(data, i)
	}
	fmt.Println(data)
	baseLen := 7
	a := len(data) / baseLen
	b := len(data) % baseLen
	for i := 0; i < a; i++ {
		del := data[i*baseLen : i*baseLen+baseLen]
		fmt.Printf("len: %v, data: %v\n", len(del), del)
	}
	if b != 0 {
		del := data[a*baseLen:]
		fmt.Printf("len: %v, data: %v\n", len(del), del)
	}
}
