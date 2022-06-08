package main

import (
	"fmt"
)

func main() {
	res := func(n1 int, n2 int) int {
		return n1 * n2
	}(2, 8)
	fmt.Println("res = ", res)
}
