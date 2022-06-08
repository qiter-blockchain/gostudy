package main

import (
	"fmt"
)

var (
	fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {
	fmt.Println(fun1(42, 44))
}
