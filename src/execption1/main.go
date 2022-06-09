package main

import (
	"errors"
	"fmt"
)

func testError(name string) (err error) {
	if name == "szc" {
		return nil
	} else {
		return errors.New("Something wrong with " + name + "...") // 定义新的错误信息
	}
}
func test2() {
	err := testError("szc")
	if err != nil {
		panic(err) // 终止程序
	}

	fmt.Println("...")
}
func main() {
	test2()
}
