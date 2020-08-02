package main

import (
	"encoding/json"
	"fmt"
)

type aa struct {
	A string
	B string
	C int
}

func main() {
	a := aa{
		A: "a",
		B: "b",
		C: 1,
	}
	bytes, err := json.Marshal(a)
	if err != nil {
		fmt.Println("json.Marshal err")
		return
	}

	//bytes = bytes[5:]
	//bytes = nil
	bytes = []byte{}

	var b aa
	b.A = "qweqwe"
	b.B = "aaaaasfdfdf"
	b.C = 333
	err = json.Unmarshal(bytes, &b)
	if err != nil {
		fmt.Println("json.Unmarshal err")
		fmt.Println(a)
		fmt.Println(b)
		return
	}

	fmt.Println(a)
	fmt.Println(b)
}
