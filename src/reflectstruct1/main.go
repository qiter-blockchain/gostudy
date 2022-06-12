package main

import (
	"fmt"
	"reflect"
)

var (
	ptr    *Student_rf
	rType  reflect.Type
	rValue reflect.Value
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
	// s := Student_rf{
	// 	Name: "szc",
	// 	Age:  23,
	// }

	// reflect_struct(s)
	ptr = &Student_rf{}                // 结构体指针
	rType = reflect.TypeOf(ptr).Elem() // 结构体反射类型

	rValue = reflect.New(rType) // 由结构体反射类型，获取新结构体指针反射值

	ptr = rValue.Interface().(*Student_rf) // 把指针反射值转成空接口，并进行类型断言

	rValue = rValue.Elem() // 由结构体指针反射值获取结构体反射值

	rValue.FieldByName("Name").SetString("szc") // 根据属性名，对结构体反射值设置值
	rValue.FieldByName("Age").SetInt(22)

	fmt.Println(*ptr) // 输出结果
}
