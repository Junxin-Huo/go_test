package main

var x, y int
var (  // 这种因式分解关键字的写法一般用于声明全局变量
    a int
    b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

var i int

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"

func main(){
    //g, h := 123, "hello"
    //g, h := e, f
    var g, h = e,f
    //i:=9 //编译错误
    println(x, y, a, b, c, d, e, f, g, h)

    c,d = d,c
    println(c,d)
    _,k := 10,11
    println(k)
}
