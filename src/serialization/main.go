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
	p0 := Person_ser{Name: "szc", Age: 23}

	json_bytes, error_ := json.Marshal(&p0)
	if error_ != nil {
		fmt.Println("Json error:", error_)
		return
	}

	fmt.Println(string(json_bytes))
	// {"Name":"szc","Age":23}
}
