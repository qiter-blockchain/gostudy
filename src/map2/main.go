package main

import "fmt"

func main() {
	var a map[string]string
	a = make(map[string]string, 10)
	b := make(map[string]string)
	c := map[string]string{
		"key1": "宋江",
		"key2": "无用", //, 不能省略
	}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}
