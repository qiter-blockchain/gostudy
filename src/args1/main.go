package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args {
		fmt.Println("第", (index + 1), "个参数是", arg)
	}

}
