package main

import (
	"fmt"
)

func main() {
	str := "拜仁慕尼黑来自德甲"
	for index, s := range str {
		fmt.Printf("%d --- %c ", index, s)
	}
}
