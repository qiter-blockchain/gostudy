package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println("welcome")
	}

	j := 1
	for j <= 10 {
		fmt.Println("cc")
		j++
	}
	a := 1
	for {
		if a <= 10 {
			fmt.Printf("a = %d \n", a)

		} else {
			fmt.Println("ok~~~")
			break
		}
		a++

	}

	var str string = "hello,world!"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c \n", str[i])
	}

	for index, val := range str {
		fmt.Printf("index=%d,value=%c\n", index, val)
	}

	var str1 string = "你好，世界"
	str2 := []rune(str1)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c \n", str2[i])
	}
}
