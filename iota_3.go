package main

import "fmt"
const (
    a=100 + iota
    b
    c
    d
)

func main() {
    fmt.Println("a=",a)
    fmt.Println("b=",b)
    fmt.Println("c=",c)
    fmt.Println("d=",d)
}
