package main

import "fmt"

func ChangeMap(input map[int]string) {
	input[2] = "changed"
	input[3] = "itsname"
	delete(input, 1)
	fmt.Printf("Inside input address: %p | %v\n", &input, input)
}

func main() {
	tmpmap := make(map[int]string)
	tmpmap[1] = "myname"
	tmpmap[2] = "yourname"
	fmt.Printf("Before change input address: %p | %v\n", &tmpmap, tmpmap)
	ChangeMap(tmpmap)
	fmt.Printf("After change input address: %p | %v\n", &tmpmap, tmpmap)
}
