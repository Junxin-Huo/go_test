package main

import "fmt"

func main() {
	var s *string = nil
	var i interface{} = s
	fmt.Println(s == nil)
	fmt.Println(i == nil)
}
