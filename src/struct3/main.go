package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name     string `json:"name"` // 标签
	Age      int    `json:"age"`
	Hometown string `json:"homeTown"`
}

func main() {
	person0 := Person{"szc", 23, "Washington"}

	jsonStr, _ := json.Marshal(person0)
	fmt.Println(string(jsonStr))
}
