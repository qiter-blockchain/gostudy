package main

import (
	"fmt"
)

type Person struct {
	Name     string `json:"name"` // 标签
	Age      int    `json:"age"`
	Hometown string `json:"homeTown"`
}

func (p Person) test() {
	fmt.Println("name:", p.Name)
}
func main() {
	person0 := Person{"Bob", 23, "California"}

	person2 := new(Person)
	(*person2).Name = "Jason"
	(*person2).Age = 24

	person0.test()
	(*person2).test()
}
