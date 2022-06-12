package main

import (
	"fmt"
	"structextends/model"
)

func main() {
	graduate0 := &model.Graduate{}
	graduate0.Name = "szc"
	graduate0.Age = 23
	graduate0.Major = "software"

	fmt.Println(graduate0.GetAge())
	fmt.Println(graduate0.GetMajor())
}
