package main

import "fmt"

func main() {
	m1 := make(map[string]string) // map[键类型] 值类型

	m1["name"] = "Jason" // 键值对赋值
	m1["age"] = "23"
	fmt.Println(m1)

	//按键取值
	fmt.Println(m1["name"])         // songzeceng
	fmt.Println(m1["gender"])       // 空字符串
	fmt.Println(m1["gender"] == "") // true

	//删除某值
	delete(m1, "age") // 如果不存在age键，则也不会报错
	//如果需要清空映射，直接分配新的内存就行
	fmt.Println(m1["age"]) // songzeceng

	for k, v := range m1 {
		fmt.Println(k, "--", v)
	}
}
