package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	i := 1

	bytes, error_ := json.Marshal(&i)
	if error_ != nil {
		fmt.Println("Json error:", error_)
		return
	}

	fmt.Println(string(bytes))
	// 1
}
