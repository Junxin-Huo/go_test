package main

import (
	"fmt"
	"time"
)

func main() {
	a := time.Now().Unix()
	fmt.Println(a)
	fmt.Println(time.Now())
	fmt.Println(time.Now().Local())
	fmt.Println(time.Unix(a, 0))
	fmt.Println(time.Unix(a, 0).String())
	//fmt.Println(math.Pi)

	zone, _ := time.Now().Zone()
	fmt.Println(zone)
}
