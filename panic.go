package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var user = os.Getenv("USER_")
	go func() {
		defer func() {
			fmt.Println("defer here")
		}()

		if user == "" {
			panic("should set user env.")
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Printf("get result %d\r\n", 1)
}