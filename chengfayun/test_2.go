package main

import (
	"fmt"
)

func charToInt(c byte) int {
	if c-'0' < 10 {
		return int(c - '0')
	}
	return int(c-'a') + 10
}

func main() {
	fromBase := 0
	number := ""
	toBase := 0

	fmt.Scan(&fromBase, &number, &toBase)

	sum := 0
	numberByte := []byte(number)
	for _, p := range numberByte {
		sum = sum*fromBase + charToInt(p)
	}
	//fmt.Println(sum)

	//fmt.Println(charToInt('0'))
	//fmt.Println(charToInt('9'))
	//fmt.Println(charToInt('a'))
	//fmt.Println(charToInt('f'))

	var ansNum []byte
	for {
		a := sum / toBase
		b := sum % toBase

		ansNum = append(ansNum, byte(b))
		if a == 0 {
			break
		}
		sum = a
	}
	//fmt.Println(ansNum)

	for i := len(ansNum); i > 0; i-- {
		fmt.Print(ansNum[i-1])
	}
	fmt.Print("\n")

	for i := len(ansNum); i > 0; i-- {
		fmt.Printf("%x", ansNum[i-1])
	}
}
