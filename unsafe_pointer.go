package main

import (
	"fmt"
	"unsafe"
)

func f1()  {
	//i:= 10
	//ip:=&i
	//
	//var fp *float64 = (*float64)(ip)
	//
	//fmt.Println(fp)
}

func f2()  {
	i:= 10
	ip:=&i

	var fp *float64 = (*float64)(unsafe.Pointer(ip))

	*fp = *fp * 3

	fmt.Println(i)
}

func f3() {
	u:=new(user)
	fmt.Println(*u)

	pName:=(*string)(unsafe.Pointer(u))
	*pName="张三"

	pAge:=(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(u))+unsafe.Offsetof(u.age)))
	*pAge = 20

	//temp:=uintptr(unsafe.Pointer(u))+unsafe.Offsetof(u.age)
	//pAge:=(*int)(unsafe.Pointer(temp))
	//*pAge = 20

	fmt.Println(*u)
}

func main() {
	f1()
	f2()
	f3()
}

type user struct {
	name string
	age int
}
