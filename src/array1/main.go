package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

	var nums [4]int = [4]int{1, 2, 3, 4}
	fmt.Println("均值为", nums)

	var nums1 = [4]int{1, 2, 3, 4}
	fmt.Println("均值为", nums1)

	var nums3 = [...]int{1, 2, 3, 4} // 自行判断长度，中括号里...一个不能少
	fmt.Println("均值为", nums3)

	var num4 = [...]int{1: 3, 0: 4, 2: 5} // 指定索引和值
	fmt.Println("均值为", num4)
}
