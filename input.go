package main

import "fmt"

func main() {
	//var buffer [512]byte
	//
	//n, err := os.Stdin.Read(buffer[:])
	//if err != nil {
	//
	//	fmt.Println("read error:", err)
	//	return
	//
	//}
	//
	//
	//fmt.Println("count:", n, ", msg:", string(buffer[:]))


	//ch := make(chan int, 5)
	////<- chan
	//
	//ch <- 11
	//ch <- 12
	//ch <- 34
	//
	//
	//
	//y, oky := <- ch
	//fmt.Println(y, oky)
	//
	//close(ch)
	//
	//for x := range ch {
	//	fmt.Println(x)
	//}
	//
	//x, ok := <- ch
	//fmt.Println(x, ok)



	var c1, c2, c3 chan int
	var i1, i2 int
	i2 = 1
	//c1 <- -1
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3):  // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}

}