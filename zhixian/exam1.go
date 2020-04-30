package main

import "fmt"

func update(data []int)  {
	max := 0
	index := 0
	for k, v := range data {
		if v > max {
			max = v
			index = k
		}
	}
	for k := range data {
		if k == index {
			data[k] -= 3
		} else {
			data[k] += 1
		}
	}
}

func main() {
	project := []int{10, 7, 5, 4}
	count := 10*12
	for i := 1; i<count; i++ {
		update(project)
		fmt.Println(project)
	}
}
