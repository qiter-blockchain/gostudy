package main

import (
	"fmt"
	"strconv"
	"time"
)

func test_r() {
	for i := 0; i < 10; i++ {
		fmt.Println("test_r test.......", strconv.Itoa(i+1))
		time.Sleep(2 * time.Second) // 休眠2秒
	}
}

func main() {
	go test_r() // 开启一个协程
	for i := 0; i < 10; i++ {
		fmt.Println("main test.......", strconv.Itoa(i+1))
		time.Sleep(time.Second)
	}
}
