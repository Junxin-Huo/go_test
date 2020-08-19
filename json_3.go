package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	s := "{\"kkk\":\"vvv\"}"
	res := make(map[string]interface{})
	err := json.Unmarshal([]byte(s), &res)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
