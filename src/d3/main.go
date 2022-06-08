package main

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break // break  默认会跳出最近的for循环
			}
			fmt.Println("j=", j)
		}
	}
}
