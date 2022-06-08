package main

import "fmt"

func main() {
	var a int = 4
	var b int32
	var c float32
	var ptr *int

	fmt.Printf("a 变量类型为 = %T\n", a) //输出变量类型%T
	fmt.Printf("b 变量类型为 = %T\n", b)
	fmt.Printf("c 变量类型为 = %T\n", c)

	ptr = &a
	fmt.Printf("a 的内存地址为 = %p", ptr) //go里面的内存块地址通常都是用十六进制表示的，因此输出：0x10414020a
	fmt.Printf("*ptr 为 %d\n", *ptr)  //这是个指向a的内存地址的指针，因此输出：4

}
