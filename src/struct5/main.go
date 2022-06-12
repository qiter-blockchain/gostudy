package main

import (
	"fmt"
)

type C struct {
	n int
}

type A struct {
	n int
}

type B struct {
	A // 匿名结构体，父类
	n int
	c C // 有名结构体，成员变量
}

func (a *A) test() {
	fmt.Println("A...")
}

func (b *B) test() {
	fmt.Println("B...")
}

func (c *C) test() {
	fmt.Println("C...")
}

func main() {
	var b B
	b.n = 10
	b.A.n = 21
	b.c.n = 31 // 显式访问成员变量的属性

	fmt.Println(b.n)
	fmt.Println(b.A.n)
	fmt.Println(b)
	fmt.Println(b.c.n)

	b.test()
	b.A.test()
	b.c.test() // 显式调用成员变量的方法
}
