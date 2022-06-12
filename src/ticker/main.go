package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	counter := 1
	for _ = range ticker.C {
		fmt.Println("ticker 1") //每秒执行一次
		counter++
		if counter >= 5 {
			break
		}
	}
	ticker.Stop() //停止
}
