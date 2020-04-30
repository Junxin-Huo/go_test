package main
import "fmt"
func main(){
    nums := []int{1,2,3,4}
    for i,num := range nums {
        fmt.Printf("索引是%d,长度是%d\n",i, num)
    }
}
