package main

import "fmt"

type ICalculate interface { // type 接口名 interface
	add()
	sub()
}
type B struct {
}

type D struct {
}

func (b B) add() {
	fmt.Println("B..add")
}

func (b B) sub() {
	fmt.Println("B..sub")
}

func (d D) add() {
	fmt.Println("D..add")
}

func (d D) sub() {
	fmt.Println("D..sub")
}

type E struct {
}

func main() {
	b0 := B{}

	var ic ICalculate
	ic = b0
	ic.add()
}
