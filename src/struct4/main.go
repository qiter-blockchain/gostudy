package main

import (
	"fmt"
)

type A struct {
	n int
}

type B struct {
	A
	n int
}

func (a *A) test() {
	fmt.Println("A...")
}

func (b *B) test() {
	fmt.Println("B...")
}

func main() {
	var b B
	b.n = 10
	b.A.n = 21

	fmt.Println(b.n)
	fmt.Println(b.A.n) // 显式调用
	fmt.Println(b)

	b.test()
	b.A.test()
}
