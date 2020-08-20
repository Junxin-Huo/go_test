package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	sum := sha256.Sum256([]byte("1234"))
	sum = sha256.Sum256(sum[:])
	fmt.Printf("sum: %x\n", sum)

	checkSum := []byte{0, 0}
	fmt.Printf("checkSum:%x\n", checkSum)

	for i := 0; i < len(sum); i += 2 {
		fmt.Printf("%x\n", sum[i:i+2])
		checkSum[0] ^= sum[i]
		checkSum[1] ^= sum[i+1]
	}

	fmt.Printf("checkSum: %x\n", checkSum)
	fmt.Printf("finallySign: %x\n", append(sum[:], checkSum...))
}
