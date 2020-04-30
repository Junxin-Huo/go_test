package main

import "fmt"

func swap(x, y string) (string, string) {
    return y, x
}

func swap2(x, y string) () {
    x = x + "111"
    y = y + "222"
}

func swap3(x, y *string) () {
    *x = *x + "---"
    *y = *y + "+++"
}

func main() {
    var a, b = "Mahesh", "Kumar"
    fmt.Println(a, b)

    fmt.Println("swap:")
    a, b = swap(a, b)
    fmt.Println(a, b)

    fmt.Println("swap:")
    swap(a, b)
    fmt.Println(a, b)

    fmt.Println("swap2:")
    swap2(a, b)
    fmt.Println(a, b)

    fmt.Println("swap3:")
    swap3(&a, &b)
    fmt.Println(a, b)
}
