package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type aa struct {
	p reflect.Type
}

type bb struct {
	s string
}

func main() {
	//v := "hello world"
	//fmt.Println(typeof(v))
	//typeTest(v, reflect.TypeOf(v))
	v2 := bb{
		s: "hello world",
	}
	fmt.Println(typeof(v2))
	typeTest(v2, reflect.TypeOf(v2))
}

func typeof(value interface{}) string {
	return reflect.TypeOf(value).String()
}

func typeTest(value interface{}, p reflect.Type)  {
	fmt.Println(p.String())
	//fmt.Println(value.(p))
	//new()
	//switch value.(type) {
	//case p:
	//	fmt.Println(p.String())
	//case int:
	//	fmt.Println("int")
	//}
	//value.(string)

	switch p.Kind() {
	case reflect.String:
		fmt.Println("typeTest", p.String())
	}

	obj := reflect.New(p)
	data, err := json.Marshal(value)
	if err != nil {
		fmt.Println("json.Marshal", err)
	}
	err = json.Unmarshal(data, obj.Interface())
	if err != nil {
		fmt.Println("json.Unmarshal", err)
	}
	inter := obj.Elem().Interface().(bb)
	fmt.Println(obj, inter)
	t := obj.Elem().FieldByName("main.bb")
	fmt.Println("ttt", t)
}
