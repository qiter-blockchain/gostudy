package main

import "fmt"

type Person struct {
	Name     string
	Age      int
	Hometown string
}

func main() {
	person0 := Person{"Jason", 23, "Washington"}

	fmt.Println(person0)

	person1 := person0
	person1.Age = 21
	fmt.Println(person0) // person0的age依旧是23
	fmt.Println(person1) // person0的age依旧是23

	person2 := new(Person) // 指针声明方式1
	(*person2).Name = "Jason"
	(*person2).Age = 24

	fmt.Println(*person2) // 没有赋值的字段默认为0值

	person3 := &Person{"Mike", 20, "London"} // 指针声明方式2
	fmt.Println(*person3)
}
