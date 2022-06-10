package main

import (
	"fmt"
)

func modify(array *[6]float64) {
	array[0] += 5
}

func main() {
	//1.
	array0 := [...]int{1, 2, 3, 4, 5, 6}
	slice := array0[1:4] // 切片  array0[1: 4] 表示slice引用array0数组起始下标为1，最后下标为4(不包含4)

	slice[0] = 7
	fmt.Println(array0[1]) // 7

	//3.
	slice2 := []int{1, 2, 4}
	fmt.Println(slice2[1]) // 7
}
