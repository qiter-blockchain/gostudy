package main

import "fmt"

func main() {
	str := "hello@world!"
	slece := str[:]
	fmt.Println("slece=", slece)
	// string 是不可变的，不能通过str[0] ='z' 方式来修改字符串
	// 修改字符串，可以将 string -> []byte  或者 []rune  -> 修改  ->重写转成string
	upstr := []byte(str)

	// 可以处理数字和英文还有符号，不能处理中文。原因：一个汉字占3个字节
	upstr[5] = '-'
	str = string(upstr)
	fmt.Println("str=", str)

	// 处理中文 转成 []rune 即可
	chstr := []rune(str)
	chstr[5] = '一'
	chstr1 := string(chstr)
	fmt.Println("chstr=", chstr1)
}
