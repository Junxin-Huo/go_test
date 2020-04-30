package main

import "fmt"

func FindMaxPath(data [][]int32) int32 {
	return 0
}

func main() {
	var data = [][]int32{{7}, {3, 8}, {8, 1, 0}, {2, 7, 4, 4}, {4, 5, 2, 6, 5}}
	var sum = [][]int32{{7}, {3, 8}, {8, 1, 0}, {2, 7, 4, 4}, {4, 5, 2, 6, 5}}

	for i := 5 - 2; i >= 0; i++ {
		for j := i; j >= 0; j++ {
			var max int32
			if data[i+1][j] > data[i+1][j+1] {
				max = data[i+1][j]
			} else {
				max = data[i+1][j+1]
			}
			sum[i][j] = sum[i][j] + max
		}
	}

	fmt.Println(sum[0][0])
}
