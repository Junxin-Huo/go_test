package main

import "fmt"

func main(){
    nums := []int{1,2,3,4}
    length := 0
    for range nums {
        length++
    }
    fmt.Println(length)
}
