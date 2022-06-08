package main

import (
	"fmt"
)

func main() {
	res := func(n1 int, n2 int) (int, int) {
		return n2, n1
	}
	n1 := 10
	n2 := 20
	n1, n2 = res(n1, n2)
	fmt.Println("res = ", n1, n2)
}
