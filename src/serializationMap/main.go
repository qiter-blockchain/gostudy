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
	var a map[string]interface{} // 键是空接口，表示能接收任意类型
	a = make(map[string]interface{})

	a["name"] = "szc"
	a["age"] = 21

	bytes, error_ := json.Marshal(&a)
	if error_ != nil {
		fmt.Println("Json error:", error_)
		return
	}

	fmt.Println(string(bytes))
	// {"age":23,"name":"szc"}
}
