package main

import "fmt"

func main() {
	rp()
	cp()
}

func rp() {
	dict := map[string]int{"key1": 1, "key2": 2}
	value, ok := dict["key1"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("key1 不存在")
	}
}
func cp() {
	dict := map[string]int{"key1": 1, "key2": 2}
	if value, ok := dict["key2"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("key1 不存在")
	}
}
