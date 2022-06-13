package main

import "fmt"

func main() {
	var sum float64 = 0.0
	for i := 1; i < 10; i++ {
		var score float64
		fmt.Println("welcome")
		fmt.Scanln(&score)
		sum += score
	}
	fmt.Printf("avg: %v", sum/9)

}
