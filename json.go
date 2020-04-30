package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	arr := []string{"hello", "apple", "python", "golang", "base", "peach", "pear"}
	fmt.Println(arr)
	lang, err := json.Marshal(arr)
	if err == nil {
		fmt.Println("================array 到 json str==")
		fmt.Println(string(lang))
	}

	//json 到 []string
	var wo []string
	if err := json.Unmarshal(lang, &wo); err == nil {
		fmt.Println("================json 到 []string==")
		fmt.Println(wo)
	}

	str := "[\"iost\",\"liketeryfox\",\"ContractFvXunxWKXe6qast1dvmBD6okbFPqk3QkqNjyMuh92RnG\",\"35\",\"iost.endless.game\"]"
	fmt.Println(str)
	fmt.Println([]byte(str))

	var aa []string
	bb, _ := json.Marshal(str)
	if err := json.Unmarshal(bb, &aa); err == nil {
		fmt.Println("================json 到 []string== bb")
		fmt.Println(aa)
	} else {
		fmt.Println(err)
	}

	if err := json.Unmarshal([]byte(str), &aa); err == nil {
		fmt.Println("================json 到 []string== []byte(str)")
		fmt.Println(aa)
	} else {
		fmt.Println(err)
	}
}
