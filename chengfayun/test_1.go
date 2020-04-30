package main

import (
	"fmt"
	"strings"
)

func main() {
	absPath := ""
	fmt.Scan(&absPath)
	//fmt.Printf("%s\n", absPath)

	subPath := strings.Split(absPath, "/")
	//fmt.Print(subPath)

	var stack []string
	for _, p := range(subPath) {
		//fmt.Println(p)
		if p == "." || p == "" {
			continue
		}
		if p == ".." {
			stack = stack[0: len(stack)-1]
		} else {
			stack = append(stack, p)
		}
	}

	//fmt.Println(stack)

	var simPath string = ""
	for _, p := range(stack) {
		//fmt.Println(p)
		simPath += "/" + p
	}
	if simPath == "" {
		simPath = "/"
	}

	fmt.Println(simPath)
}
