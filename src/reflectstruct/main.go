package main

import (
	"fmt"
	"reflect"
)

type Student_rf struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Student_rf_0 struct {
	Name  string
	Sex   string
	Score int
	Age   int
}

func (s Student_rf) Show(name string, age int) {
	fmt.Println(name, " -- ", age)
}

func (s Student_rf) GetAge() int {
	return s.Age
}
func reflect_struct(n interface{}) {
	rType := reflect.TypeOf(n)
	rValue := reflect.ValueOf(n)

	iv := rValue.Interface()

	fmt.Println("rType = ", rType, ", iv = ", iv)
	// rType =  main.Student_rf , iv =  {szc 23}
	fmt.Printf("Type of iv = %T\n", iv)
	// Type of iv = main.Student_rf

	// 类型断言
	switch iv.(type) {
	case Student_rf:
		student := iv.(Student_rf)
		fmt.Println(student.Name, ", ", student.Age)
	case Student_rf_0:
		student := iv.(Student_rf_0)
		fmt.Println(student.Name, ", ", student.Age)
	}
	fmt.Println("")           // 打印方法地址
	num := rValue.NumMethod() // 获取方法数量
	for i := 0; i < num; i++ {
		method := rValue.Method(i)
		fmt.Println(method) // 打印方法地址
	}

	var params []reflect.Value
	params = append(params, reflect.ValueOf("szc"))
	params = append(params, reflect.ValueOf(21))
	rValue.Method(1).Call(params) // 调用方法，传入参数

	fmt.Println("...")

	res := rValue.Method(0).Call(nil) // 调用方法，接收返回值
	fmt.Println(res[0].Int())

}

func main() {
	s := Student_rf{
		Name: "szc",
		Age:  23,
	}

	reflect_struct(s)

	fmt.Println("")
	// 修改结构体字段的值，就要和修改普通类型变量的值一样，获取地址的引用
	rValue := reflect.ValueOf(&s)
	rValue.Elem().Field(0).SetString("szc")
	rValue.Elem().Field(1).SetInt(24)

	fmt.Println(s)
}
