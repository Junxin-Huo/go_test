package main

import "fmt"

func main() {
	s := fmt.Sprint("1234%v www", "pp")
	fmt.Println(s)
}
