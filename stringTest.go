package main

import "fmt"

func main() {
	a := ""
	b := "0x"
	c := "0x12345678"
	d := "12345678"
	fmt.Println(len(a), "|", len(b), "|", len(c), "|", len(d))
	fmt.Println("c[0:2]", c[0:2])
	fmt.Println("d[0:2]: ", d[0:2])

	fmt.Println("c[2:]", c[2:])
	fmt.Println("c[2:10]", c[2:10])
	//fmt.Println("c[2:11]", c[2:11])    // panic: runtime error: slice bounds out of range


	input := "0xa9059cbb"
	input += "0000000000000000000000003715c698addd4896482b1e92a2c41c9ed8be3210"
	input += "0000000000000000000000000000000000000000000000001bc16d674ec80000"
	fmt.Println("len(input)", len(input), input)
	input = input[2:]
	fmt.Println("len(input)", len(input), input)

	funSel := input[:8]
	input = input[8:]

	param1 := input[:64]
	param1 = param1[24:]  // 24 = 64 - 40
	input = input[64:]

	param2 := input[:64]
	param2 = param2[24:]  // 24 = 64 - 40
	input = input[64:]

	fmt.Println("len(funSel)", len(funSel), funSel)
	fmt.Println("len(param1)", len(param1), param1)
	fmt.Println("len(param2)", len(param2), param2)
	fmt.Println("len(input)", len(input), input)
}
