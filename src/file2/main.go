package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("D:/output.txt")
	if err != nil {
		fmt.Println("open file error = ", err)
		return
	}

	fmt.Println("file = ", *file)
	bytes, err1 := ioutil.ReadFile("D:/output.txt")
	if err1 != nil {
		fmt.Println("open file error = ", err1)
		return
	}

	fmt.Println(string(bytes))
	err = file.Close()
	if err != nil {
		fmt.Println("close file error = ", err)
	}

}
