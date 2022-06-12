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
	fmt.Println("E add ICalculate")
	ic.add()
}

func (e *E) sub(ic ICalculate) {
	fmt.Println("E sub ICalculate")
	ic.sub()
}
func main() {
	b0 := B{}
	d0 := D{}
	e0 := E{}
	b0.add()
	d0.add()
	fmt.Println("")
	e0.add(b0)
	e0.add(d0)
	e0.sub(b0)
	e0.sub(d0)
}
