package main

import (
	"fmt"
	"math/rand"
	"time"
)

func modify(array *[6]float64) {
	array[0] += 5
}

func main() {
	var hens [6]float64
	total := 0.0

	rand.Seed(time.Now().Unix())
	for i := 0; i < len(hens); i++ {
		hens[i] = rand.Float64()*20 + 5
		fmt.Println("第", (i + 1), " 个数是", hens[i])

		total += hens[i]
	}

	fmt.Println("均值为", (total / float64(len(hens))))

	modify(&hens)

	total = 0
	for i := 0; i < len(hens); i++ {
		fmt.Println("第", (i + 1), " 个数是", hens[i])
		total += hens[i]
	}

	fmt.Println("均值为", (total / float64(len(hens))))
}
