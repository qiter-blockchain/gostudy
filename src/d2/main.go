package main

import "fmt"

func main() {
	var j string
	fmt.Scanln(&j) // Scanln读取一行
	fmt.Println("j = ", j)

	var i string
	var m float32
	var n bool

	fmt.Scanf("%d%f%s%t", &i, &m, &j, &n)
	fmt.Println("i = ", i, "j = ", j, "m = ", m, "n = ", n)
}
