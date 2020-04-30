package main

import (
	"fmt"
	"io/ioutil"
	"bytes"
)

func read1(buf *bytes.Buffer) {
	fmt.Println("0: ", len(buf.Bytes()))
	ioutil.ReadAll(buf)
	fmt.Println("1: ", len(buf.Bytes()))
}

func peek(buf *bytes.Buffer, b []byte) (int, error) {
	buf2 := bytes.NewBuffer(buf.Bytes())
	return buf2.Read(b)
}

func read2(buf *bytes.Buffer)  {
	fmt.Printf("Len: %d, Content: %s\n", buf.Len(), buf)

	fmt.Println("\nPeeking...")
	data := make([]byte, 4)
	peek(buf, data)
	fmt.Printf("Peeked: %s\n", data)
	fmt.Printf("Len: %d, Content: %s\n", buf.Len(), buf)

	fmt.Println("\nReading...")
	data = make([]byte, buf.Len())
	buf.Read(data)
	fmt.Printf("Read: %s\n", data)
	fmt.Printf("Len: %d, Content: %s\n", buf.Len(), buf)
}

func main() {
	buf := &bytes.Buffer{}
	buf.WriteString("Hello")
	read2(buf)
}
