package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	typ := reflect.StructOf([]reflect.StructField{
		{
			Name: "Height",
			Type: reflect.TypeOf(float64(0)),
			Tag:  `json:"Height"`,
		},
		{
			Name: "Age",
			Type: reflect.TypeOf(int(0)),
			Tag:  `json:"Age"`,
		},
		{
			Name: "Data",
			Type: reflect.StructOf([]reflect.StructField{
				{
					Name: "Height",
					Type: reflect.TypeOf(float64(0)),
					Tag:  `json:"Height"`,
				},
				{
					Name: "Age",
					Type: reflect.TypeOf(int(0)),
					Tag:  `json:"Age"`,
				},
			}),
			Tag: `json:"Data"`,
		},
	})

	v := reflect.New(typ).Elem()
	v.Field(0).SetFloat(0.4)
	v.Field(1).SetInt(2)
	s := v.Addr().Interface()

	w := new(bytes.Buffer)
	if err := json.NewEncoder(w).Encode(s); err != nil {
		panic(err)
	}

	fmt.Printf("value: %+v\n", s)
	fmt.Printf("json:  %s", w.Bytes())

	r := bytes.NewReader([]byte(`{"height":1.5,"age":10, "Data":{"Height":100, "Age":333}}`))
	if err := json.NewDecoder(r).Decode(s); err != nil {
		panic(err)
	}
	fmt.Printf("value: %+v\n", s)

}
