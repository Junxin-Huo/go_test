package main

import (
	"fmt"
	"regexp"
)

func main() {
	reg, err := regexp.Compile("^[a-z0-9_]+$")
	if err != nil {
		fmt.Println(err, reg)
		return
	}
	fmt.Println(reg.MatchString("12_3asdz_xc"))
	fmt.Println(reg.MatchString("ABC"))
	fmt.Println(reg.MatchString("我你他"))
	fmt.Println(reg.MatchString("_"))
}
