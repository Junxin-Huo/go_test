package main

import (
	"crypto/rand"
)

var (
	src = make([]byte, 100)
	dst = make([]byte, 100)
)

func main() {
	rand.Read(src)
	// 效率低
	dst = append(dst, src...)

	rand.Read(src)
	// 效率更高
	copy(dst, src)
}
