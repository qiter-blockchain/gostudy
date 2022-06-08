package main

import (
	"fmt"
)

func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
}

func main() {
	f := AddUpper()
	fmt.Println(f(1)) // 11
	fmt.Println(f(3)) // 14
	fmt.Println(f(3)) // 17

	g := AddUpper()
	fmt.Println(g(4)) // 14
}
