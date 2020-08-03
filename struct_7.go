package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1/test")
	{
		v1.GET("", Test)
	}
}

func Test(c *gin.Context) {
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
			Name: "Name",
			Type: reflect.TypeOf(string("")),
			Tag:  `json:"Name"`,
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
	s := v.Addr().Interface()

	err := c.BindJSON(s)
	if err != nil {
		fmt.Println(c, "bind json error %v", err)
		c.String(http.StatusBadRequest, "bind json error %v", err)
		return
	}

	v.Field(0).SetFloat(v.Field(0).Float() * 10)
	v.Field(1).SetInt(v.Field(1).Int() * 100)
	v.Field(2).SetString(v.Field(2).String() + "===+++")
	v.Field(3).Field(0).SetFloat(v.Field(3).Field(0).Float() + 1)
	v.Field(3).Field(1).SetInt(v.Field(3).Field(1).Int() + 2)

	bytes, err := json.Marshal(s)
	if err != nil {
		fmt.Println(c, "json marshal error %v", err)
		c.String(http.StatusBadRequest, "bind json error %v", err)
		return
	}
	c.JSON(200, string(bytes))
}
