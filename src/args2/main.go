package main

import (
	"flag"
	"fmt"
)

func main() {
	var s0 string
	var s1 string
	var i0 int

	flag.StringVar(&s0, "u", "", "字符串参数1") // 接收字符串参数，参数列表：参数接收地址，参数名，默认值， 参数说明
	flag.StringVar(&s1, "p", "", "字符串参数2")
	flag.IntVar(&i0, "i", 0, "整型参数1")

	flag.Parse() // 开始解析必须调用

	fmt.Println("s0 = ", s0, ", s1 = ", s1, ", i0 = ", i0)
}
