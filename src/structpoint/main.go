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
func (p *Person) String() string {
	return fmt.Sprintf("name:%v\tage:%v\thometown:%v\tscore:%v", p.Name, p.Age, p.Hometown, p.score)
}

func main() {
	m0 := make(map[string]int)
	m0["China"] = 80
	person0 := Person{"Mike", 23, "Manchester", m0}

	person1 := person0
	person1.Age = 21

	person2 := new(Person)
	(*person2).Name = "Jason"
	(*person2).Age = 24
	m2 := make(map[string]int)
	m2["Math"] = 90
	(*person2).score = m2

	fmt.Println(&person0)
	fmt.Println(person0)
	fmt.Println(person2)
	fmt.Println(*person2)
}
