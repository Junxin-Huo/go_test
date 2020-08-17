package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var count int32 = 100
	var offset int32 = 100
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().Unix())
		offset = rand.Int31n(count)
		fmt.Println(offset)
	}

	fmt.Println("###############")
	for i := 0; i < 10; i++ {
		offset = rand.Int31n(count)
		fmt.Println(offset)
	}
}
