package main

import "fmt"

func main()  {
	a := []int{22, 10, 100, 23}
	M := 60
	b := map[int]int{}
	for i := 0; i<len(a) ;i++ {
		if M % a[i] == 0 {
			//b = append(b, a[i]) //1,2,3,30
			_, ok := b[M/a[i]]
			if ok {
				println(a[i], " * ", M/a[i], " = ", M)
			} else {
				b[a[i]] = M/a[i]
			}
		}
	}



	c := []int{}
	for i := 1; i<len(a); i++ {
		c = append(c, a[i] - a[i-1])
	}

	sum := 0
	max := 0
	//day1 := -1
	//day2 := -1
	for i := 0; i<len(c); i++ {
		sum = sum + c[i]
		if sum < 0 {
			sum = 0
		}

		if sum > max {
			max = sum
		}
	}
	fmt.Println(max)



	dayBegin := 0
	sum = 0
	max = 0
	for i := 1; i<len(a); i++ {
		sum = a[i] - a[dayBegin]
		if sum < 0 {
			dayBegin = i
			sum = 0
		}

		if sum > max {
			max = sum
		}
	}
	fmt.Println(max)
}


















