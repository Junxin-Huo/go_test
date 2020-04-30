package main

import "fmt"

func charCheck(c uint8) (uint8, error) {
	if c <= '9' && c >= '0' {
		return c - '0', nil
	}
	if c <= 'z' && c >= 'a' {
		return (c - 'a') % 9 + 1, nil
	}
	return 0, nil
}

func main()  {
	code := "12345abcde"

	if len(code) != 16 {
		fmt.Println("error")
		return
	}

	// 是否为奇数
	flag := true
	// 奇数和
	sum1 := 0
	// 偶数和
	sum2 := 0
	for i := len(code) - 1; i >= 0; i-- {
		a, err := charCheck(code[i])
		if err != nil {
			fmt.Println("error")
			return
		}

		if flag {
			sum1 += int(a)
		} else {
			tmp := int(a) * 2
			if tmp > 10 {
				tmp -= 9
			}
			sum2 += tmp
		}
		flag = !flag
	}

	if (sum1+sum2)%10 == 0 {
		fmt.Println("ok")
		return
	}
}