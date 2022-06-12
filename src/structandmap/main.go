package main

import (
	"fmt"
)

type Person struct {
	Name     string
	Age      int
	Hometown string
	score    map[string]int
}

func (p Person) test() {
	p.Age += 1
	p.score["China"] += 1
}

// 为了能使方法里的修改更高效地同步到外面，声明方法时一般会绑定结构体指针
func (p *Person) test_1(n int) string {
	(*p).score["China"] += n
	(*p).Age -= n

	return "succeed"
}

// 如果要实现类似java里的toString，我们可以对指定数据类型绑定String()方法，返回string
func (p Person) String() string {
	return fmt.Sprintf("name:%v\tage:%v\thometown:%v\tscore:%v", p.Name, p.Age, p.Hometown, p.score)
}

type integer int // 要先定义别名

func (i *integer) test(n int) {
	*i += integer(n) // int和integer虽然只是别名关系，但依旧不是同一个类型
}

func main() {
	m0 := make(map[string]int)
	m0["China"] = 80
	person0 := Person{"szc", 23, "Henan Anyang", m0}

	person2 := new(Person)
	(*person2).Name = "Jason"
	(*person2).Age = 24
	m2 := make(map[string]int)
	m2["Math"] = 90
	(*person2).score = m2
	// 会得到以下输出，age没有变，但映射属性却发生了改变
	person0.test()
	fmt.Println(person0)
	(*person2).test()
	fmt.Println(*person2)
	// 调用时，还是可以直接使用变量名调用方法，而不必取址
	fmt.Println(m0, "\n", m2)
	fmt.Println("")
	(&person0).test_1(5)
	fmt.Println(person0)
	person0.test_1(5)
	fmt.Println(person0)

	var num integer
	num = 8
	num.test(6)
	fmt.Println(num)
}
