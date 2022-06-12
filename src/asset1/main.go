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

func (e *E) add(ic ICalculate) { // 接口是引用类型，所以这里传递的是变量的引用
	ic.add()
}

func (e *E) sub(ic ICalculate) {
	ic.sub()
}
func main() {
	b0 := B{}
	d0 := D{}
	e0 := E{}

	e0.add(b0)
	e0.add(d0)
	e0.sub(b0)
	e0.sub(d0)
	var cals []ICalculate
	cals = append(cals, b0)
	cals = append(cals, d0)

	fmt.Println(cals)
	b1, succeed := cals[0].(B) // 使用方法：待断言变量.(断言类型)
	if succeed {
		fmt.Println("convert success")
	} else {
		fmt.Println("convert fail")
	}
	c1, succeed := cals[1].(B)
	if succeed {
		fmt.Println("convert success")
	} else {
		fmt.Println("convert fail")
	}

	fmt.Println(b1)
	fmt.Println(c1)
}
