package main

import (
	"fmt"
	"strconv"
	"time"
)

func test_r_0() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test_r_0 发生错误:", err)
		}
	}()
	var map0 map[int]string
	map0[0] = "szc"
}
func test_r() {
	for i := 0; i < 10; i++ {
		fmt.Println("test_r test.......", strconv.Itoa(i+1))
		time.Sleep(500 * time.Millisecond)
	}
}
func main() {

	go test_r()
	go test_r_0()

	time.Sleep(time.Second * 10)
}
