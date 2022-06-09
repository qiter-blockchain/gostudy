package main

import (
	"fmt"
)

func test() {
	defer func() {
		err := recover() // 捕获异常
		if err != nil {
			fmt.Println("err:", err) // 输出异常
		}
	}()
	n1 := 1
	n2 := 0
	fmt.Println("res:", n1/n2)
}

func main() {
	test()
}
