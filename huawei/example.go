// 本题为考试单行多行输入输出规范示例，无需提交，不计分。
package main

import (
	"fmt"
)
func main() {
	a:=0
	b:=0
	for {
		n, _ := fmt.Scan(&a,&b)
		if n == 0 {
			break
		} else {
			fmt.Printf("%d\n",a+b)
		}
	}
}