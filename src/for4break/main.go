package main

import "fmt"

func main() {
lable1:
	for i := 1; i < 10; i++ {
		// lable2:
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				fmt.Println("k", k)
				break lable1
			}
			fmt.Println("j")
		}
		fmt.Println("i")
	}

}
