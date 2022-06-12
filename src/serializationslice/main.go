package main

import (
	"encoding/json"
	"fmt"
)

type Person_ser struct {
	Name string `json:"name"` // 标签
	Age  int    `json:"age"`
}

func main() {
	var slice []map[string]interface{}
	slice = make([]map[string]interface{}, 0)

	for i := 0; i < 5; i++ {
		var a map[string]interface{}
		a = make(map[string]interface{})

		a["name"] = "szc" + fmt.Sprintf("%d", i)
		a["age"] = 23

		slice = append(slice, a)
	}

	bytes, error_ := json.Marshal(&slice)
	if error_ != nil {
		fmt.Println("Json error:", error_)
		return
	}

	fmt.Println(string(bytes))
	// [{"age":23,"name":"szc0"},{"age":23,"name":"szc1"},{"age":23,"name":"szc2"},{"age":23,"name":"szc3"},{"age":23,"name":"szc4"}]
}
