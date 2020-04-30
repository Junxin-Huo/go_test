package main

import (
	"fmt"
)

func main() {
	data := ""
	n, _ := fmt.Scan(&data)
	if n == 0 {
		return
	}

	//for k, v := range data {
	//	fmt.Println(k, v)
	//}

	mark := map[int32]int{}
	count := 0
	for _, value := range data {
		_, ok := mark[value]
		//fmt.Println(value, ok)
		if !ok {
			count++
			mark[value] = 1
		}
	}

	fmt.Println(count)
}
