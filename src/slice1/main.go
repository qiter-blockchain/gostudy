package main

import "fmt"

func modify(array *[6]float64) {
	array[0] += 5
}

func main() {
	slice1 := make([]int, 10)
	array0 := [...]int{1, 2, 3, 4, 5, 6}
	slice := array0[1:4]
	copy(slice1, slice) // 参数列表：dest、src
	fmt.Println(slice1)
	slice1[len(slice1)-1] = -1
	fmt.Println(slice) // 原切片不变
}
