package main

import "fmt"

func main() {
    var  ptr *int

    fmt.Printf("ptr 的值为 : %x\n", ptr  )
    if ptr == nil {
        fmt.Printf("nil\n")
    } else {
        fmt.Printf("!nil\n")
    }
}
