package main

import "fmt"

type Point struct {
	x, y int
}
type Rect struct {
	leftUp, rightDown Point
}
type Rect_ struct {
	leftUp, rightDown *Point
}

func main() {
	rect0 := Rect{Point{1, 2}, Point{3, 4}}
	fmt.Println(rect0)
	fmt.Printf("%p \n", &rect0)
	fmt.Println(&rect0.leftUp.x, &rect0.leftUp.y, &rect0.rightDown.x, rect0.rightDown.y)

}
