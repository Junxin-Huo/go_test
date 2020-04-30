package main

import (
	"fmt"
	"math"
)

func CheckEqual(a, b string) bool {
	return false
}

func main() {
	a := ""
	b := ""

	n, _ := fmt.Scan(&a, &b)
	if n == 0 {
		return
	}

	sum_a := 0.0
	dot := false
	count := 0
	dec := 0.0
	for _, v := range a {
		//fmt.Println(k, v)
		if v == '.' {
			dot = true
		} else {
			if dot {
				count ++
				//if v != '0' {
				//	sum_a = sum_a + (float64(v) - '0')/math.Pow(10, float64(count))
				//}
				dec = dec +  (float64(v) - '0')
			} else {
				sum_a = sum_a*10 + (float64(v) - '0')
			}
		}
	}
	for i := 0; i<count; i++ {
		dec /= 10
	}
	sum_a += dec

	sum_b := 0.0
	dot = false
	count = 0
	for _, v := range b {
		//fmt.Println(k, v)
		if v == '.' {
			dot = true
		} else {
			if dot {
				count ++
				if v != '0' {
					sum_b = sum_b + (float64(v) - '0')/math.Pow(10, float64(count))
				}
			} else {
				sum_b = sum_b*10 + (float64(v) - '0')
			}
		}
	}

	fmt.Println(sum_a, sum_b)
}
