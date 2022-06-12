package main

import (
	"fmt"
	"reflect"
)

func reflect_base(n int) {
	rTyp := reflect.TypeOf(n)                   // 获取反射类型
	fmt.Println("rType = ", rTyp)               // int
	fmt.Println("rType`s name = ", rTyp.Name()) // int

	rVal := reflect.ValueOf(n)                                  // 获取反射值
	fmt.Printf("rValue = %v, rValue`s type = %T\n", rVal, rVal) // 100, reflect.Value

	n1 := 2 + rVal.Int() // 获取反射值持有的整型值
	fmt.Println("n1 = ", n1)

	iV := rVal.Interface() // 反射值转换成空接口
	num := iV.(int)        // 类型断言
	fmt.Println("num = ", num)
}
func main() {
	reflect_base(1)
}
