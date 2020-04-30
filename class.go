package main

import "fmt"

//创建一个Birds类
type Birds struct {
	Name string  //具名字段（聚合）
}
//Birds类有三个方法
func (bird *Birds)HasFoot()  {
	fmt.Println(bird.Name," has feet!")
}

func (bird *Birds)HasEye()  {
	fmt.Println(bird.Name," has eye!")
}

func (bird *Birds)HasFlying()  {
	fmt.Println(bird.Name,"FATHER: can flying!")
}
//创建一个Ostrich类，并继承 Birds
type Ostrich struct {
	Birds //匿名字段（嵌入）
	Wings string  //具名字段（聚合）
}
//重写Base的Flying方法
func (ostrich *Ostrich)HasFlying()  {
	ostrich.Birds.HasFoot()   //显示调用父类Birds的方法
	ostrich.Birds.HasEye()     //显示调用父类Birds的方法
	ostrich.HasFoot()            //隐式调用父类Birds的方法
	ostrich.HasEye()              //隐式调用父类Birds的方法
	fmt.Println(ostrich.Name,"SON: can`t fly")

	fmt.Println(ostrich.Name,"SON: 111")
	ostrich.Birds.HasFlying()

	//fmt.Println(ostrich.Name,"SON: 222")
	//ostrich.HasFlying()
}

func main() {
	ostrich := new(Ostrich)
	ostrich.Name = "鸵鸟"
	ostrich.HasFlying() //调用子类的方法
}