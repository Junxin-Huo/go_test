package main

import (
	"fmt"
	"strings"
)

func main() {
	s := strings.Split("abc,abc", ",")
	fmt.Println(s, len(s))
	s = strings.Split("abc,abc,", ",")
	fmt.Println(s, len(s))
}
