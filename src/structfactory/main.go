package main

import (
	"fmt"
	"structfactory/model"
)

func main() {
	student0 := model.CreateStudent("Jason", 23) // 调用公有方法，获得指针对象
	fmt.Println(*student0)
	fmt.Println(student0.GetAge())
}
