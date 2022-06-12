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
	str := "{\"Name\":\"szc\",\"Age\":23}"
	var p0 Person_ser

	err := json.Unmarshal([]byte(str), &p0)

	if err != nil {
		fmt.Println("Json error:", err)
		return
	}

	fmt.Println(p0)
	// {szc 23}
}
