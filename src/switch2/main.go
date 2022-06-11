package main

import (
	"fmt"
)

func main() {
	var score float64
	fmt.Println("input a score:")
	fmt.Scanln(&score)

	switch int(score / 60) {
	case 1:
		fmt.Println("A")
	case 0:
		fmt.Println("b")
	default:
		fmt.Println("Unknow")
	}

}
