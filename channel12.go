package main

import "fmt"

func main() {
	jobs := make(chan int, 10)
	for i := 0; i < 10; i++ {
		jobs <- i
	}
	close(jobs)

	a, b := <- jobs
	fmt.Println(a, b)

	for t := range jobs {
		fmt.Println(t)
	}
	fmt.Println(<-jobs, "=====")

	a, b = <- jobs
	fmt.Println(a, b)
}
