package main

import "fmt"

type A struct {
	A1 string
	A2 int
}

type B struct {
	B1 string
	B2 int
	B3 *A
}

type C struct {
	C1 string
	C2 int
	C3 *B
}

func (a *A) Set(m string, n int) {
	a.A1 = m
	a.A2 = n
}

func (c *C) Get() *A {
	return c.C3.B3
}

func (c *C) GetB() *B {
	return c.C3
}

func (c *C) Set(a *A) {
	c.C3.B3 = a
}

func main() {
	c := C{
		C1: "c1",
		C2: 1,
		C3: &B{
			B1: "B1",
			B2: 2,
			B3: &A{
				A1: "B1",
				A2: 2,
			},
		},
	}

	a1 := c.Get()
	fmt.Println(a1)
	fmt.Println("------------------")

	c.C3.B3.Set("change", 666)

	fmt.Println(a1)
	a2 := c.Get()
	fmt.Println(a2)

	b2 := c.GetB()
	fmt.Println(b2, b2.B3)
	fmt.Println("------------------")

	c.Set(&A{
		A1: "happy",
		A2: 2333,
	})

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(b2, b2.B3)
	a3 := c.Get()
	fmt.Println(a3)
	fmt.Println("------------------")
}
