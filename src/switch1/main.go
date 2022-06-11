package main

import (
	"fmt"
)

func main() {
	var char byte
	fmt.Println("input a char:")
	fmt.Scanf("%c", &char)

	switch char {
	case 'a':
		fmt.Println("A")
	case 'b':
		fmt.Println("b")
	case 'c':
		fmt.Println("C")
	default:
		fmt.Println("default")
	}

}
