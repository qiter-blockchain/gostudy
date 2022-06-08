package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(generateRandom(time.Now().Unix(), 100))
	fmt.Println(sum(100))
}
func generateRandom(time int64, _range int) int {
	rand.Seed(time)
	return rand.Intn(_range)
}
func sum(args ...int) int { //支持0-多个参数
	return args[0]
}
func init() {
	fmt.Println("init variable_advanced..")
}
